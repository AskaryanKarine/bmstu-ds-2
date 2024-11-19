package clients

import (
	"encoding/json"
	"fmt"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/models"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type ReservationClient struct {
	client  httpClient
	baseUrl string
}

func NewReservationClient(client httpClient, baseUrl string) *ReservationClient {
	return &ReservationClient{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (r *ReservationClient) GetReservationByUUID(username, uuid string) (models.ExtendedReservationResponse, error) {
	urlReq := fmt.Sprintf("%s/%s/%s", r.baseUrl, "reservations", uuid)
	req, err := http.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return models.ExtendedReservationResponse{}, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("X-User-Name", username)

	resp, err := r.client.Do(req)
	if err != nil {
		return models.ExtendedReservationResponse{}, fmt.Errorf("failed to make request: %w", err)
	}

	if resp == nil {
		return models.ExtendedReservationResponse{}, models.EmptyResponseError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ExtendedReservationResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var respModel models.ExtendedReservationResponse
		if err := json.Unmarshal(body, &respModel); err != nil {
			return models.ExtendedReservationResponse{}, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		return respModel, nil
	case http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError:
		var respErr models.ErrorResponse
		if err := json.Unmarshal(body, &respErr); err != nil {
			return models.ExtendedReservationResponse{}, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		respErr.StatusCode = resp.StatusCode
		return models.ExtendedReservationResponse{}, respErr
	default:
		return models.ExtendedReservationResponse{}, models.ErrorResponse{
			StatusCode: resp.StatusCode,
			Message:    models.UndefinedResponseCodeError.Error(),
		}
	}
}

func (r *ReservationClient) GetReservationsByUser(username string) ([]models.ExtendedReservationResponse, error) {
	urlReq := fmt.Sprintf("%s/%s", r.baseUrl, "reservations")
	req, err := http.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("X-User-Name", username)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	if resp == nil {
		return nil, models.EmptyResponseError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var respModel []models.ExtendedReservationResponse
		if err := json.Unmarshal(body, &respModel); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		return respModel, nil
	case http.StatusBadRequest, http.StatusNotFound, http.StatusInternalServerError:
		var respErr models.ErrorResponse
		if err := json.Unmarshal(body, &respErr); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		respErr.StatusCode = resp.StatusCode
		return nil, respErr
	default:
		return nil, models.ErrorResponse{
			StatusCode: resp.StatusCode,
			Message:    models.UndefinedResponseCodeError.Error(),
		}
	}
}

func (r *ReservationClient) CancelReservation(username, uuid string) error {
	urlReq := fmt.Sprintf("%s/%s/%s", r.baseUrl, "reservations", uuid)
	req, err := http.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("X-User-Name", username)

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	if resp == nil {
		return models.EmptyResponseError
	}
	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError:
		var respErr models.ErrorResponse
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}
		if err := json.Unmarshal(body, &respErr); err != nil {
			return fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		resp.Body.Close()
		respErr.StatusCode = resp.StatusCode
		return respErr
	default:
		return models.ErrorResponse{
			StatusCode: resp.StatusCode,
			Message:    models.UndefinedResponseCodeError.Error(),
		}
	}
}

func (r *ReservationClient) GetHotels(page, size int) (models.PaginationResponse, error) {
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("size", strconv.Itoa(size))
	urlReq := fmt.Sprintf("%s/%s?%s", r.baseUrl, "hotels", params.Encode())
	req, err := http.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return models.PaginationResponse{}, fmt.Errorf("failed to build request: %w", err)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return models.PaginationResponse{}, fmt.Errorf("failed to make request: %w", err)
	}
	if resp == nil {
		return models.PaginationResponse{}, models.EmptyResponseError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.PaginationResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		var respModel models.PaginationResponse
		if err := json.Unmarshal(body, &respModel); err != nil {
			return models.PaginationResponse{}, fmt.Errorf("failed to parse response body: %w", err)
		}
		return respModel, nil
	case http.StatusBadRequest:
		var valErr models.ValidationErrorResponse
		if err := json.Unmarshal(body, &valErr); err != nil {
			return models.PaginationResponse{}, fmt.Errorf("failed to parse response body: %w", err)
		}
		return models.PaginationResponse{}, valErr
	case http.StatusInternalServerError:
		var respErr models.ErrorResponse
		if err := json.Unmarshal(body, &respErr); err != nil {
			return models.PaginationResponse{}, fmt.Errorf("failed to parse response body: %w", err)
		}
		respErr.StatusCode = resp.StatusCode
		return models.PaginationResponse{}, respErr
	default:
		return models.PaginationResponse{}, models.ErrorResponse{
			StatusCode: resp.StatusCode,
			Message:    models.UndefinedResponseCodeError.Error(),
		}
	}

}
