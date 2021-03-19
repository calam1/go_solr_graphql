package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/calam1/orderhistory/graph/generated"
	"github.com/calam1/orderhistory/graph/model"
	"github.com/mitchellh/mapstructure"
	"github.com/vanng822/go-solr/solr"
)

func (r *queryResolver) OrderHistory(ctx context.Context, account *model.UserAccountIds) ([]*model.OrderHistory, error) {
	si, _ := solr.NewSolrInterface("https://solr.com:8983/solr", "orders")
	si.SetBasicAuth("solr-user", "solr-pwd")

	query := solr.NewQuery()
	// todo: if we uncomment group and group sort we have to change the struct since the format is different, for now just leave it, just a learning thing now
	query.AddParam("group.sort", "orderLineNumber")
	query.AddParam("qt", "orders")
	query.AddParam("group.main", "false")
	query.AddParam("group.format", "grouped")
	query.AddParam("group.ngroups", "true")
	query.AddParam("group.field", "orderNumber")
	query.FieldList("*")
	query.Rows(10)
	query.Start(0)
	query.Sort("orderCreatedDateTime desc")
	query.Q("accountNumber:" + account.AccountNumber)
	query.FilterQuery("contactId:" + account.ContactId)

	s := si.Search(query)
	result, err := s.Result(nil)
	if err != nil {
		fmt.Printf("error in solr result %v", err)
	}

	var orderHistories []*model.OrderHistory
	for _, x := range result.Results.Docs {
		var orderHistory model.OrderHistory
		mapstructure.Decode(x, &orderHistory)
		orderHistories = append(orderHistories, &orderHistory)
	}

	log.Printf("result %+v", orderHistories)

	return orderHistories, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
