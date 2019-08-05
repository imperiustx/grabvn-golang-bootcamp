package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/learn-grpc/feedback/feedbackpb"
	"google.golang.org/grpc"
)

var c feedbackpb.FeedbackServiceClient

func main() {
	fmt.Println("feedback client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c = feedbackpb.NewFeedbackServiceClient(cc)

	for {
		fmt.Println("hello what do you want to do?\n1: Add feedback\n2: Get passenger feedback ID\n3: Get booking code\n4: Delete feedback")

		var chooseOneNum int
		_, err := fmt.Scan(&chooseOneNum)
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		var passengerID int32
		var bookingCode string
		var feedback string
		switch chooseOneNum {
		case 1:
			fmt.Println("Input your PassengerID here")
			fmt.Scan(&passengerID)
			fmt.Println("Input your booking code here")
			fmt.Scan(&bookingCode)
			fmt.Println("Input your feedback here")
			fmt.Scan(&feedback)
			addFeedback(ctx, &feedbackpb.PassengerFeedback{
				PassengerID: passengerID,
				BookingCode: bookingCode,
				Feedback:    feedback,
			})
		case 2:
			var passengerID int32
			fmt.Println("Please input PassengerID")
			fmt.Scan(&passengerID)
			getFeedbackByPassagerID(ctx, passengerID)

		case 3:
			var bookingCode string
			fmt.Println("Please input BookingCode")
			fmt.Scan(&bookingCode)
			getFeedbackByBookingCode(ctx, bookingCode)

		case 4:
			var passengerID int32
			fmt.Println("Please input PassengerID")
			fmt.Scan(&passengerID)
			deleteFeedbackByPassagerID(ctx, passengerID)
		}
	}

}

func addFeedback(ctx context.Context, passengerFeedback *feedbackpb.PassengerFeedback) {
	require, err := c.AddPassengerFeedback(ctx, &feedbackpb.AddPassengerFeedbackRequest{NewPassengerFeedback: passengerFeedback})
	if err != nil {
		log.Printf("could not add the feedback by params %v because of %v", &passengerFeedback, err)
	} else {
		log.Printf("Add Feedback response: %s", require.GetNewFeedback())
	}
}

func getFeedbackByPassagerID(ctx context.Context, passengerID int32) {
	require, err := c.GetPassengerFeedbackID(ctx, &feedbackpb.GetPassengerFeedbackIDRequest{PassengerID: passengerID})
	if err != nil {
		log.Printf("could not get feedbacks by passager id %v because of %v", passengerID, err)
	} else {
		log.Printf("GetFeedbackByPassagerID response: %s", require.GetFeedbackID())
	}
}

func getFeedbackByBookingCode(ctx context.Context, bookingCode string) {
	require, err := c.GetBookingCode(ctx, &feedbackpb.GetBookingCodeRequest{BookingCode: bookingCode})
	if err != nil {
		log.Printf("could not get feedbacks by booking code %v becasue of %v", bookingCode, err)
	} else {
		log.Printf("GetFeedbackByBookingCode response: %s", require.GetPassengerFeedback())
	}
}

func deleteFeedbackByPassagerID(ctx context.Context, passengerID int32) {
	require, err := c.DeleteByPassengerFeedbackID(ctx, &feedbackpb.DeleteByPassengerFeedbackIDRequest{PassengerID: passengerID})
	if err != nil {
		log.Printf("could not delete feedbacks by passager id %v because of %v", passengerID, err)
	} else {
		log.Printf("DeleteFeedbackByPassagerID response: %s", require.GetPassengerFeedback())
	}
}
