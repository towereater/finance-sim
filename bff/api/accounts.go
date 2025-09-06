package api

import (
	"bff/config"
	"bff/model"
	"bff/service"
	"encoding/json"
	"fmt"
	acc "mainframe-lib/account/model"
	sacc "mainframe-lib/account/service"
	com "mainframe-lib/common/config"
	"net/http"
	"strconv"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	queryParams := r.URL.Query()
	var err error

	from := queryParams.Get(string(config.ContextFrom))
	if from != "" && len(from) != 24 {
		fmt.Printf("Invalid %s parameter\n", string(config.ContextFrom))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := 50
	if queryParams.Has(string(config.ContextLimit)) {
		limit, err = strconv.Atoi(queryParams.Get(string(config.ContextLimit)))

		if err != nil {
			fmt.Printf("Invalid %s parameter\n", string(config.ContextLimit))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if limit > 50 {
			limit = 50
		}
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	auth := r.Context().Value(com.ContextAuth).(string)
	userId := r.Context().Value(config.ContextUserId).(string)

	// Build the filter
	var filter acc.Account
	filter.Id.Service = queryParams.Get(string(config.ContextService))
	filter.Owner = userId

	// Get all accounts
	accounts, err := sacc.GetAccounts(cfg.Services.Accounts, cfg.Services.Timeout, auth, filter, from, limit)
	if err != nil {
		fmt.Printf("Error while searching accounts with filter %+v: %s\n",
			filter, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get account additional data
	var accountInfos model.GetAccountsOutput
	for _, a := range accounts {
		if a.Id.Service == "CK" {
			ckAccount, err := service.GetCheckingAccountInfo(cfg, auth, a.Id.Account)
			if err != nil {
				fmt.Printf("Error while searching checking account %+v: %s\n",
					a.Id, err.Error())
				continue
			}

			accountInfos.Accounts = append(accountInfos.Accounts, ckAccount)
		} else {
			accountInfos.Accounts = append(accountInfos.Accounts, model.AccountInfo{
				AccountId: a.Id,
			})
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accountInfos)
}
