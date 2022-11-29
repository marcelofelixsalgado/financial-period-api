package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Return a JSON response to request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Rerturn an error in JSON format
func JSONErrorByCode(w http.ResponseWriter, errorCode ErrorCode) {
	message := NewResponseMessage()
	message.AddMessageByErrorCode(errorCode)
	//jsonMessage, _ := message.GetJsonMessage()
	JSON(w, message.GetMessage().HttpStatusCode, message.GetMessage())
	// w.WriteHeader(message.GetMessage().HttpStatusCode)
	// w.Write(jsonMessage)
}

// func JSONErrorByIssue(w http.ResponseWriter, issue Issue) {
// 	message := NewResponseMessage()
// 	message.AddMessageByIssue(issue)
// 	//jsonMessage, _ := message.GetJsonMessage()
// 	JSON(w, message.GetMessage().HttpStatusCode, message.GetMessage())
// 	w.WriteHeader(message.GetMessage().HttpStatusCode)
// 	// w.Write(jsonMessage)
// }
