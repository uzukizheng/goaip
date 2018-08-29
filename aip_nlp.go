package goaip

import "strings"

type AIPNlp struct {
	AIPBase
}

func MakeAIPNlp(appID string, key string, secret string) *AIPNlp {
	an := new(AIPNlp)
	an.AppID = appID
	an.Key = key
	an.Secret = secret
	return an
}

//词法分析
func (an *AIPNlp) Lexer(txt string) (string, error) {
	an.mutex.Lock()
	defer an.mutex.Unlock()
	var data = make(map[string]interface{})
	data["text"] = strings.Trim(txt, " ")
	if r, e := an.Request(LEXER_URL, data, 10); e != nil {
		return "", e
	} else {
		return r, nil
	}
}

func (an *AIPNlp) Simnet(lhs string, rhs string) (string, error) {
	an.mutex.Lock()
	defer an.mutex.Unlock()
	var data = make(map[string]interface{})
	data["text_1"] = strings.Trim(lhs, " ")
	data["text_2"] = strings.Trim(rhs, " ")
	return an.Request(SIMNET_URL, data, 10)
}

func (an *AIPNlp) DepParse(txt string) (string, error) {
	an.mutex.Lock()
	defer an.mutex.Unlock()
	var data = make(map[string]interface{})
	data["mode"] = 0
	data["text"] = txt
	return an.Request(DEP_PARSER_URL, data, 10)
}

func (an *AIPNlp) WordEmbedding(word string) (string, error) {
	an.mutex.Lock()
	defer an.mutex.Unlock()
	var data = make(map[string]interface{})
	data["word"] = word
	return an.Request(WORD_EMBEDDING_URL, data, 10)
}

func (an *AIPNlp) WordSimEmbedding(wordLhs string, wordRhs string) (string, error) {
	an.mutex.Lock()
	defer an.mutex.Unlock()
	var data = make(map[string]interface{})
	data["word_1"] = wordLhs
	data["word_2"] = wordRhs
	return an.Request(WORD_SIM_EMBEDDING_URL, data, 10)
}

func (an *AIPNlp) SentimentClassify(text string) (string, error) {
	an.mutex.Lock()
	defer an.mutex.Unlock()
	var data = make(map[string]interface{})
	data["text"] = text
	return an.Request(SENTIMENT_CLASSIFY_URL, data, 10)
}

func (an *AIPNlp) Keyword(title string, content string) (string, error) {
	an.mutex.Lock()
	defer an.mutex.Unlock()
	var data = make(map[string]interface{})
	data["title"] = title
	data["content"] = content
	return an.Request(KEYWORD_URL, data, 10)
}

func (an *AIPNlp) Topic(title string, content string) (string, error) {
	an.mutex.Lock()
	defer an.mutex.Unlock()
	var data = make(map[string]interface{})
	data["title"] = title
	data["content"] = content
	return an.Request(TOPIC_URL, data, 10)
}
