package main

import (
	"./exreader"
	pb "./passenger"
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	address = "localhost:8080"
)

var reader = exreader.ExReader{Reader: bufio.NewReader(os.Stdin)}

func selectMenu() int32 {
	fmt.Println("Actions Menu:")
	fmt.Println("	1. Add feedback")
	fmt.Println("	2. Get feedback by booking code")
	fmt.Println("	3. Get feedback by passenger id")
	fmt.Println("	4. Delete feedback by passenger id")
	fmt.Println("	5. Exit")
	fmt.Print("Please select your action: ")

	actionNum, err := reader.ReadInt32()

	if err != nil {
		fmt.Printf("Error: %v. Please press enter to return actions menu.", err)
		fmt.Scanln()

		return -1
	}

	return actionNum
}

func clearConsole() {
	var c = exec.Command("clear")
	c.Stdout = os.Stdout

	if err := c.Run(); err != nil {
		log.Fatalf("Can't not clear console %v", err)
	}
}

func getFeedbackByBookingCode(client pb.PassengerServiceClient) () {
	var err error
	var bookingCode string
	var requestData *pb.GetPassengerFeedbackRequest

	clearConsole()
	fmt.Print("Please enter booking code: ")

	bookingCode, err = reader.ReadText()

	if err != nil {
		fmt.Printf("Error: %v. Please press enter to return actions menu.", err)
		fmt.Scanln()
		return
	}

	requestData = &pb.GetPassengerFeedbackRequest{
		BookingCode: bookingCode,
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetPassengerFeedbackByBookingCode(ctx, requestData)

	if err != nil {
		log.Fatalf("Couldn't get feedback by booking code: %v", err)
	}

	if response.ResponseCode.Code != pb.SuccessResponse.Code {
		fmt.Println(response.ResponseCode.Message+".", "Please press enter to return actions menu.")
	} else {
		var feedback = response.PassengerFeedback
		fmt.Println("Results: ")
		fmt.Printf("\t Passenger ID: %d - Booking code: %s - Feedback: %s\n", feedback.PassengerID, feedback.BookingCode, feedback.Feedback)
		fmt.Println("Please press enter to return actions menu.")

	}

	fmt.Scanln()
}

func getFeedbackByPassengerID(client pb.PassengerServiceClient) () {
	var err error
	var passengerID int32
	var requestData *pb.GetPassengerFeedbackRequest

	fmt.Print("Passenger ID: ")
	passengerID, err = reader.ReadInt32()

	if err != nil {
		fmt.Printf("Error: %v. Please press enter to return actions menu.", err)
		fmt.Scanln()
		return
	}

	requestData = &pb.GetPassengerFeedbackRequest{
		PassengerID: passengerID,
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetPassengerFeedbackByPassengerId(ctx, requestData)

	if err != nil {
		log.Fatalf("Couldn't get feedback by passenger id: %v", err)
	}

	if response.ResponseCode.Code != pb.SuccessResponse.Code {
		fmt.Println(response.ResponseCode.Message+".", "Please press enter to return actions menu.")
	} else {
		var feedbacks = response.PassengerFeedbacks
		fmt.Println("Results: ")

		for _, feedback := range feedbacks {
			fmt.Printf("\t Passenger ID: %d - Booking code: %s - Feedback: %s\n", feedback.PassengerID, feedback.BookingCode, feedback.Feedback)
		}
		fmt.Println("Please press enter to return actions menu.")
	}

	fmt.Scanln()
}

func deleteFeedbackByPassengerID(client pb.PassengerServiceClient) () {
	var err error
	var passengerID int32
	var requestData *pb.DeletePassengerFeedbackRequest

	fmt.Print("Passenger ID: ")
	passengerID, err = reader.ReadInt32()

	if err != nil {
		fmt.Printf("Error: %v. Please press enter to return actions menu.", err)
		fmt.Scanln()
		return
	}

	requestData = &pb.DeletePassengerFeedbackRequest{
		PassengerID: passengerID,
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.DeletePassengerFeedbackPassengerId(ctx, requestData)

	if err != nil {
		log.Fatalf("Couldn't add feedback: %v", err)
	}

	fmt.Println(response.ResponseCode.Message+".", "Please press enter to return actions menu.")
	fmt.Scanln()
}

func addPassengerFeedback(client pb.PassengerServiceClient) () {
	var err error
	var feedback *pb.PassengerFeedback

	feedback, err = inputFeedback()

	if err != nil {
		fmt.Printf("Error: %v. Please press enter to return actions menu.", err)
		fmt.Scanln()
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.AddPassengerFeedback(ctx, feedback)

	if err != nil {
		log.Fatalf("Couldn't add feedback: %v", err)
	}

	fmt.Println(response.ResponseCode.Message+".", "Please press enter to return actions menu.")
	fmt.Scanln()
}

func inputFeedback() (*pb.PassengerFeedback, error) {
	var err error
	var passengerID int32
	var bookingCode string
	var feedback string

	clearConsole()
	fmt.Println("Please enter feedback information")

	// Get passengerID
	fmt.Print("Passenger ID: ")
	passengerID, err = reader.ReadInt32()

	if err != nil {
		return nil, err
	}

	// Get bookingCode
	fmt.Print("Booking Code: ")
	bookingCode, err = reader.ReadText()

	if err != nil {
		return nil, err
	}

	// Get feedback
	fmt.Print("Feedback: ")
	feedback, err = reader.ReadText()

	if err != nil {
		return nil, err
	}

	return &pb.PassengerFeedback{
		PassengerID: passengerID,
		BookingCode: bookingCode,
		Feedback:    feedback,
	}, err
}

func main() {
	//Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	defer conn.Close()

	var passengerServiceClient = pb.NewPassengerServiceClient(conn)

	for {
		switch action := selectMenu(); action {
		case 1:
			addPassengerFeedback(passengerServiceClient)
		case 2:
			getFeedbackByBookingCode(passengerServiceClient)
		case 3:
			getFeedbackByPassengerID(passengerServiceClient)
		case 4:
			deleteFeedbackByPassengerID(passengerServiceClient)
		case 5:
			fmt.Println("Bye bye...")
			return
		case -1:
			// Just for ignore default when select have error
		default:
			fmt.Println("Action isn't exists. Please press enter to return actions menu.")
			fmt.Scanln()
		}

		clearConsole()
	}
}
