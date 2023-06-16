package ksyun

import (
	"github.com/KscSDK/ksc-sdk-go/service/bws"
	"github.com/KscSDK/ksc-sdk-go/service/ebs"
	"github.com/KscSDK/ksc-sdk-go/service/eip"
	"github.com/KscSDK/ksc-sdk-go/service/epc"
	"github.com/KscSDK/ksc-sdk-go/service/iam"
	"github.com/KscSDK/ksc-sdk-go/service/kcm"
	"github.com/KscSDK/ksc-sdk-go/service/kcsv1"
	"github.com/KscSDK/ksc-sdk-go/service/kcsv2"
	"github.com/KscSDK/ksc-sdk-go/service/kec"
	"github.com/KscSDK/ksc-sdk-go/service/knad"
	"github.com/KscSDK/ksc-sdk-go/service/krds"
	"github.com/KscSDK/ksc-sdk-go/service/mongodb"
	"github.com/KscSDK/ksc-sdk-go/service/rabbitmq"
	"github.com/KscSDK/ksc-sdk-go/service/sks"
	"github.com/KscSDK/ksc-sdk-go/service/slb"
	"github.com/KscSDK/ksc-sdk-go/service/sqlserver"
	"github.com/KscSDK/ksc-sdk-go/service/tag"
	"github.com/KscSDK/ksc-sdk-go/service/tagv2"
	"github.com/KscSDK/ksc-sdk-go/service/vpc"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
	//"github.com/KscSDK/ksc-sdk-go/service/knad"
)

type KsyunClient struct {
	region        string               `json:"region,omitempty"`
	dryRun        bool                 `json:"dry_run,omitempty"`
	eipconn       *eip.Eip             `json:"eipconn,omitempty"`
	slbconn       *slb.Slb             `json:"slbconn,omitempty"`
	vpcconn       *vpc.Vpc             `json:"vpcconn,omitempty"`
	kecconn       *kec.Kec             `json:"kecconn,omitempty"`
	sqlserverconn *sqlserver.Sqlserver `json:"sqlserverconn,omitempty"`
	krdsconn      *krds.Krds           `json:"krdsconn,omitempty"`
	kcmconn       *kcm.Kcm             `json:"kcmconn,omitempty"`
	sksconn       *sks.Sks             `json:"sksconn,omitempty"`
	kcsv1conn     *kcsv1.Kcsv1         `json:"kcsv_1_conn,omitempty"`
	kcsv2conn     *kcsv2.Kcsv2         `json:"kcsv_2_conn,omitempty"`
	epcconn       *epc.Epc             `json:"epcconn,omitempty"`
	ebsconn       *ebs.Ebs             `json:"ebsconn,omitempty"`
	mongodbconn   *mongodb.Mongodb     `json:"mongodbconn,omitempty"`
	ks3conn       *s3.S3               `json:"ks_3_conn,omitempty"`
	iamconn       *iam.Iam             `json:"iamconn,omitempty"`
	rabbitmqconn  *rabbitmq.Rabbitmq   `json:"rabbitmqconn,omitempty"`
	bwsconn       *bws.Bws             `json:"bwsconn,omitempty"`
	tagconn       *tagv2.Tagv2         `json:"tagconn,omitempty"`
	tagv1conn     *tag.Tag             `json:"tagv1conn,omitempty"`
	knadconn      *knad.Knad           `json:"knadconn,omitempty"`
}
