package myaws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/hacker65536/gen-aws-extend-switch-roles-config/pkg/utils"
)

var (
	region = "ap-northeast-1"
)

func GetAccountsList() {

	r := listAccounts()
	for _, v := range r {
		if v.Status == "ACTIVE" {
			fmt.Println(aws.ToString(v.Name), aws.ToString(v.Id))
		}
	}

	fmt.Println(len(r))
}

func GenConfig() {
	c := utils.NewMyColors()
	r := listAccounts()
	for _, v := range r {
		if v.Status == "ACTIVE" {
			fmt.Printf("[%s]\nrole_arn = arn:aws:iam::%s:role/OrganizationAccountAccessRole\ncolor = %s\nregion = %s\n", aws.ToString(v.Name), aws.ToString(v.Id), c.GetRandomColorUniq(), region)
		}
	}
}

func listAccounts() []types.Account {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := organizations.NewFromConfig(cfg)

	input := &organizations.ListAccountsInput{
		// valid range minimum 1 ,max 20
		MaxResults: aws.Int32(20),
	}

	a := []types.Account{}

	p := organizations.NewListAccountsPaginator(svc, input)

	for p.HasMorePages() {
		o, err := p.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("unable to get next page, %v", err)
		}

		a = append(a, o.Accounts...)
	}

	return a

}
