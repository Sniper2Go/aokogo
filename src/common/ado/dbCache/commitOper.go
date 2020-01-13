package dbCache

/*
	purpose: db commit from logic opertion. 
	date: 20200113 14:45
*/

// push model for db cache, record flush db times, at next flush db, then judge last flush or not, if yes then get data from cache.
func PushCommitModels(identify string, models ...string){
	cache := getDBCache()
	for _, m := range models {
		cache.push(identify, m)
	}
}

// check model is exist, if not then flush db direct.
func HasExistCache(identify string, model string)bool{
	cache := getDBCache()
	return cache.hasExist(identify, model)
}

// first commit all update, then pop all push models.
func PopCommitModels(identify string) {
	cache := getDBCache()
	cache.updateDB(identify)
	cache.pop(identify)
}

// offline or other operate to force commit.
func ProtectCommit(identify string){
	PopCommitModels(identify)
}