package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/learn-grpc/feedback/feedbackpb"
	"google.golang.org/grpc"
)

type server struct{}

var database []feedbackpb.PassengerFeedback

func main() {
	fmt.Println("feedback server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	s := grpc.NewServer()
	feedbackpb.RegisterFeedbackServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to served: %v", err)
	}
}

func addNewPassengerFeedback(newfb feedbackpb.PassengerFeedback) {
	database = append(database, newfb)
}

func getFeedbackByPassengerID(passID int32) (arrOfAPassengerFeedbacks []*feedbackpb.PassengerFeedback) {
	for _, feedback := range database {
		if feedback.PassengerID == passID {
			arrOfAPassengerFeedbacks = append(arrOfAPassengerFeedbacks, &feedback)
		}
	}
	return
}

func getFeedbackByBookingCode(bkcode string) *feedbackpb.PassengerFeedback {
	for _, feedback := range database {
		if feedback.BookingCode == bkcode {
			return &feedback
		}
	}
	return nil
}

func delFeedbackByPassengerID(passID int32) {
	for i, feedback := range database {
		if feedback.PassengerID == passID {
			database = database[:i+copy(database[i:], database[i+1:])]
		}
	}
	return
}

func (*server) AddPassengerFeedback(ctx context.Context, req *feedbackpb.AddPassengerFeedbackRequest) (*feedbackpb.AddPassengerFeedbackResponse, error) {
	fmt.Printf("add feedback rpc: %v", req)
	newPassangerFeedback := feedbackpb.PassengerFeedback{
		Feedback:    *req.GetNewFeedback(),
		BookingCode: *req.GetBookingCode(),
		PassengerID: *req.GetPassengerID(),
	}

	addNewPassengerFeedback(newPassangerFeedback)

	return &feedbackpb.AddPassengerFeedbackResponse{}, nil

}

func (*server) GetPassengerFeedbackID(ctx context.Context, req *feedbackpb.GetPassengerFeedbackIDRequest) (*feedbackpb.GetPassengerFeedbackIDResponse, error) {
	fmt.Printf("get passenger id rpc: %v", req)
	arrOfAPassengerFeedbacks := getFeedbackByPassengerID(req.PassengerID)

	res := &feedbackpb.GetPassengerFeedbackIDResponse{
		FeedbackID: arrOfAPassengerFeedbacks,
	}
	return res, nil
}

func (*server) GetBookingCode(ctx context.Context, req *feedbackpb.GetBookingCodeRequest) (*feedbackpb.GetBookingCodeResponse, error) {
	fmt.Printf("get booking code rpc: %v", req)
	passengerFeedback := getFeedbackByBookingCode(req.BookingCode)
	if passengerFeedback != nil {
		res := &feedbackpb.GetPassengerFeedbackIDResponse{
			PassengerFeedback: passengerFeedback,
		}
	}

	return res, nil
}

func (*server) DeleteByPassengerFeedbackID(ctx context.Context, req *feedbackpb.DeleteByPassengerFeedbackIDRequest) (*feedbackpb.DeleteByPassengerFeedbackIDResponse, error) {
	fmt.Printf("delete passenger rpc: %v", req)
	db := delFeedbackByPassengerID(req.PassengerID)
	res := &feedbackpb.DeleteByPassengerFeedbackIDResponse{
		PassengerFeedback: db,
	}
	return res, nil
}
