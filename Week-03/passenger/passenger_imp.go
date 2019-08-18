package passenger

import (
	"context"
)

type PassengerFeedbackServerImp struct {
	FeedbackMap map[string]*PassengerFeedback
}


var SuccessResponse = &ResponseCode{
	Code:0,
	Message:"Success",
}

var NotFoundFeedbackResponse = &ResponseCode{
	Code: 1,
	Message:"Not found Feedback",
}

var ExistsFeedbackResponse = &ResponseCode{
	Code: 2,
	Message:"Exists Feedback",
}

func (s *PassengerFeedbackServerImp) AddPassengerFeedback(ctx context.Context, pf *PassengerFeedback) (*AddPassengerFeedbackResponse, error)  {
	var responseCode = SuccessResponse

	if _, ok := s.FeedbackMap[pf.BookingCode]; ok{
		responseCode = ExistsFeedbackResponse
	}else{
		s.FeedbackMap[pf.BookingCode] = pf
	}

	return &AddPassengerFeedbackResponse{
		ResponseCode:responseCode,
	}, nil
}

func (s *PassengerFeedbackServerImp) GetPassengerFeedbackByPassengerId(ctx context.Context, pfr *GetPassengerFeedbackRequest) (*GetPassengerFeedbacksResponse, error) {
	var responseCode = SuccessResponse
	var feedbacks []*PassengerFeedback

	for _, v := range s.FeedbackMap{
		if v.PassengerID == pfr.PassengerID{
			feedbacks = append(feedbacks, v)
		}
	}

	if len(feedbacks) == 0{
		responseCode = NotFoundFeedbackResponse
	}

	return &GetPassengerFeedbacksResponse{
		ResponseCode: responseCode,
		PassengerFeedbacks: feedbacks,
	}, nil
}

func (s *PassengerFeedbackServerImp) GetPassengerFeedbackByBookingCode(ctx context.Context,pfr *GetPassengerFeedbackRequest) (*GetPassengerFeedbackResponse, error) {
	var responseCode = NotFoundFeedbackResponse
	var feedback *PassengerFeedback

	if v, ok := s.FeedbackMap[pfr.BookingCode]; ok{
		responseCode = SuccessResponse
		feedback = v
	}

	return &GetPassengerFeedbackResponse{
		ResponseCode: responseCode,
		PassengerFeedback: feedback,
	}, nil
}

func (s *PassengerFeedbackServerImp) DeletePassengerFeedbackPassengerId(ctx context.Context, pfr *DeletePassengerFeedbackRequest) (*DeletePassengerFeedbackResponse, error) {
	var bookingCodes []string

	for _, v := range s.FeedbackMap{
		if v.PassengerID == pfr.PassengerID{
			bookingCodes = append(bookingCodes, v.BookingCode)
			//fmt.Println("Add ", v.BookingCode)
		}
	}

	for _, bookingCode := range bookingCodes{
		delete(s.FeedbackMap, bookingCode)
	}

	return &DeletePassengerFeedbackResponse{
		ResponseCode: SuccessResponse,
	}, nil
}