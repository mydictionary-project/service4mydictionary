package service4mydictionary

import "github.com/zzc-tongji/vocabulary4mydictionary"

// ServiceInterface : service interface
type ServiceInterface interface {
	GetServiceName() string
	GetCache() *CacheStruct
	Query(vocabulary4mydictionary.VocabularyAskStruct) vocabulary4mydictionary.VocabularyAnswerStruct
}
