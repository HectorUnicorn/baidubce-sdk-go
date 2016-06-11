package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	bce "baidubce"
	"baidubce/bos"
)

var bosClient = bos.DefaultClient

func getBucketLocation() {
	option := &bce.SignOption{
		//Timestamp:                 "2015-11-20T10:00:05Z",
		ExpirationPeriodInSeconds: 1200,
		Headers: map[string]string{
			"host":                "bj.bcebos.com",
			"other":               "other",
			"x-bce-meta-data":     "meta data",
			"x-bce-meta-data-tag": "meta data tag",
			//"x-bce-date":          "2015-11-20T07:49:05Z",
			//"date": "2015-11-20T10:00:05Z",
		},
		HeadersToSign: []string{"host", "date", "other", "x-bce-meta-data", "x-bce-meta-data-tag"},
	}

	location, err := bosClient.GetBucketLocation("baidubce-sdk-go", option)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(location.LocationConstraint)
}

func listBuckets() {
	bucketSummary, err := bosClient.ListBuckets(nil)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(bucketSummary.Buckets)
}

func createBucket() {
	err := bosClient.CreateBucket("baidubce-sdk-go-create-bucket-example", nil)

	if err != nil {
		log.Println(err)
	}
}

func doesBucketExist() {
	// exists, err := bosClient.DoesBucketExist("baidubce-sdk-go-create-bucket-example", nil)
	exists, err := bosClient.DoesBucketExist("guoyao11122", nil)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(exists)
}

func deleteBucket() {
	bucketName := "baidubce-sdk-go-delete-bucket-example"
	err := bosClient.CreateBucket(bucketName, nil)

	if err != nil {
		log.Println(err)
	} else {
		err := bosClient.DeleteBucket(bucketName, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func setBucketPrivate() {
	bucketName := "baidubce-sdk-go"
	err := bosClient.SetBucketPrivate(bucketName, nil)

	if err != nil {
		log.Println(err)
	}
}

func setBucketPublicRead() {
	bucketName := "baidubce-sdk-go"
	err := bosClient.SetBucketPublicRead(bucketName, nil)

	if err != nil {
		log.Println(err)
	}
}

func setBucketPublicReadWrite() {
	bucketName := "baidubce-sdk-go"
	err := bosClient.SetBucketPublicReadWrite(bucketName, nil)

	if err != nil {
		log.Println(err)
	}
}

func getBucketAcl() {
	bucketName := "baidubce-sdk-go"
	bucketAcl, err := bosClient.GetBucketAcl(bucketName, nil)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(bucketAcl.Owner)

	for _, accessControl := range bucketAcl.AccessControlList {
		for _, grantee := range accessControl.Grantee {
			log.Println(grantee.Id)
		}
		for _, permission := range accessControl.Permission {
			log.Println(permission)
		}
	}
}

func setBucketAcl() {
	bucketName := "baidubce-sdk-go"
	bucketAcl := bos.BucketAcl{
		AccessControlList: []bos.Grant{
			bos.Grant{
				Grantee: []bos.BucketGrantee{
					bos.BucketGrantee{Id: "ef5a4b19192f4931adcf0e12f82795e2"},
				},
				Permission: []string{"FULL_CONTROL"},
			},
		},
	}

	err := bosClient.SetBucketAcl(bucketName, bucketAcl, nil)

	if err != nil {
		log.Println(err)
	}
}

func putObject() {
	bucketName := "baidubce-sdk-go"

	objectKey := "put-object-from-string.txt"
	str := "Hello World 你好"
	putObjectResponse, bceError := bosClient.PutObject(bucketName, objectKey, str, nil, nil)

	if bceError != nil {
		log.Println(bceError)
	} else {
		log.Println(putObjectResponse)
		log.Println(putObjectResponse.GetETag())
	}

	return

	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	filePath := path.Join(pwd, "baidubce", "examples", "baidubce-sdk-go-test.pdf")

	objectKey = "pdf/put-object-from-bytes.pdf"
	byteArray, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println(err)
	} else {
		putObjectResponse, bceError = bosClient.PutObject(bucketName, objectKey, byteArray, nil, nil)

		if bceError != nil {
			log.Println(bceError)
		} else {
			log.Println(putObjectResponse)
			log.Println(putObjectResponse.GetETag())
		}
	}

	objectKey = "pdf/put-object-from-file.pdf"
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Println(err)
	} else {
		putObjectResponse, bceError = bosClient.PutObject(bucketName, objectKey, file, nil, nil)

		if bceError != nil {
			log.Println(bceError)
		} else {
			log.Println(putObjectResponse)
			log.Println(putObjectResponse.GetETag())
		}
	}
}

func main() {
	putObject()
	return
	getBucketAcl()
	setBucketAcl()
	getBucketLocation()
	listBuckets()
	createBucket()
	doesBucketExist()
	deleteBucket()
	setBucketPublicReadWrite()
	setBucketPublicRead()
	setBucketPrivate()
}
