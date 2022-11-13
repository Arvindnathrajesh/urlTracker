package services

import (
	"../domain"
	"../utils"
	"github.com/mpvl/unique"
)

func CreateLinkData(linkData *domain.LinkData) (*domain.LinkData, *utils.RestErr) {
	linkData, restErr := domain.CreateLinkData(linkData)
	if restErr != nil {
		return nil, restErr
	}
	return linkData, nil
}

func UrlClicked() (bool, *utils.RestErr) {

	linksData := *FindLinksData()

	var shortUrls []string
	for _, linkData := range linksData {
		shortUrls = append(shortUrls, linkData.ShortUrl)
	}

	unique.Strings(&shortUrls)
	for i := 0; i < len(shortUrls); i++ {
		boolValue, restErr := domain.UrlClicked(shortUrls[i])
		if restErr != nil {
			return boolValue, restErr
		}
	}

	return true, nil
}

func Redirect(shortUrl string) (*domain.LinkData, *utils.RestErr) {

	linkData, nil := FindLinkData(shortUrl)
	return linkData, nil
}

func FindLinkData(shortUrl string) (*domain.LinkData, *utils.RestErr) {
	linkData, restErr := domain.FindLinkData(shortUrl)
	if restErr != nil {
		return nil, restErr
	}
	return linkData, nil
}

func FindLinksData() *[]domain.LinkData {
	linksData, restErr := domain.FindLinksData()
	if restErr != nil {
		return nil
	}
	return linksData
}
