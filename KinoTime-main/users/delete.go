package users

import (
	"PracticeServer/db"
	"PracticeServer/logging"
	"PracticeServer/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := db.DB.Delete(&models.User{}, id).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Logger.Error("Failed to delete user: ", zap.Error(err))
		response := Response{Status: "fail", Message: "Failed to delete user: " + err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	logging.Logger.Info("User deleted successfully", zap.String("id", id), zap.String("func", "DeleteUsers"))
	response := Response{Status: "success", Message: "User deleted successfully"}
	json.NewEncoder(w).Encode(response)
}
