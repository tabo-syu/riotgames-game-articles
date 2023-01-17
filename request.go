package riotgamesgamearticles

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/exp/slices"
)

// domain
const (
	lolSiteDomain      = "www.leagueoflegends.com"
	valorantSiteDomain = "playvalorant.com"
)

const (
	EnUs = "en-us"
	EnGb = "en-gb"
	DeDe = "de-de"
	EsEs = "es-es"
	FrFr = "fr-fr"
	ItIt = "it-it"
	EnPl = "en-pl"
	PlPl = "pl-pl"
	ElGr = "el-gr"
	RoRo = "ro-ro"
	HuHu = "hu-hu"
	CsCz = "cs-cz"
	EsMx = "es-mx"
	PtBr = "pt-br"
	JaJp = "ja-jp"
	RuRu = "ru-ru"
	TrTr = "tr-tr"
	EnAu = "en-au"
	KoKr = "ko-kr"
	IdId = "id-id"
	ThTh = "th-th"
	ViVn = "vi-vn"
	ZhTw = "zh-tw"
	ArAe = "ar-ae"
)

// supported locales
var (
	LOLLocales      = []string{EnUs, EnGb, DeDe, EsEs, FrFr, ItIt, EnPl, PlPl, ElGr, RoRo, HuHu, CsCz, EsMx, PtBr, JaJp, RuRu, TrTr, EnAu, KoKr}
	ValorantLocales = []string{EnUs, EnGb, DeDe, EsEs, FrFr, ItIt, PlPl, RuRu, TrTr, EsMx, IdId, KoKr, PtBr, JaJp, ThTh, ViVn, ZhTw, ArAe}
)

type websiteRequest struct {
	Req *http.Request
}

func newWebsiteRequest(ctx context.Context, url string) (*websiteRequest, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate request: %w", err)
	}

	req.Header.Add("User-Agent", "RiotGames Game Articles (https://github.com/tabo-syu/riotgames-game-articles)")

	return &websiteRequest{Req: req}, nil
}

func newLOLWebsiteRequest(ctx context.Context, locale string) (*websiteRequest, error) {
	if !slices.Contains(LOLLocales, locale) {
		return nil, fmt.Errorf("invalid locale specified for %s: %s", lolSiteDomain, locale)
	}

	req, err := newWebsiteRequest(ctx, fmt.Sprintf("https://%s/page-data/%s/latest-news/page-data.json", lolSiteDomain, locale))
	if err != nil {
		return nil, fmt.Errorf("failed to generate LOLWebsiteRequest: %w", err)
	}

	return req, nil
}

func newValorantWebsiteRequest(ctx context.Context, locale string) (*websiteRequest, error) {
	if !slices.Contains(ValorantLocales, locale) {
		return nil, fmt.Errorf("invalid locale specified for %s: %s", valorantSiteDomain, locale)
	}

	req, err := newWebsiteRequest(ctx, fmt.Sprintf("https://%s/page-data/%s/news/page-data.json", valorantSiteDomain, locale))
	if err != nil {
		return nil, fmt.Errorf("failed to generate NewValorantWebsiteRequest: %w", err)
	}

	return req, nil
}
