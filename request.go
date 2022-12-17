package riotgamespatchnotes

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/exp/slices"
)

// domain
const (
	lolSiteDomain = "www.leagueoflegends.com"
)

// supported locales
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
)

var (
	LOLLocales = []string{EnUs, EnGb, DeDe, EsEs, FrFr, ItIt, EnPl, PlPl, ElGr, RoRo, HuHu, CsCz, EsMx, PtBr, JaJp, RuRu, TrTr, EnAu, KoKr}
)

type GameUpdatesRequest struct {
	Req *http.Request
}

func NewGameUpdatesRequest(ctx context.Context, domain string, locale string) (*GameUpdatesRequest, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://%s/page-data/%s/news/game-updates/page-data.json", domain, locale),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate request: %w", err)
	}

	return &GameUpdatesRequest{Req: req}, nil
}

func NewLOLGameUpdatesRequest(ctx context.Context, locale string) (*GameUpdatesRequest, error) {
	if !slices.Contains(LOLLocales, locale) {
		return nil, fmt.Errorf("invalid locale specified for %s: %s", lolSiteDomain, locale)
	}

	return NewGameUpdatesRequest(ctx, lolSiteDomain, locale)
}