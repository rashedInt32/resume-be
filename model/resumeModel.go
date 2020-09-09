// Package model provides data structure
package model

type body struct {
	Background string `json:"background" bson:"background"`
}

type display struct {
	FontFamily string `json:"fontFamily" bson:"fontFamily"`
}

type name struct {
	Fontsize       int    `json:"fontSize" bson:"fontSize"`
	FontFamily     string `json:"fontFamily" bson:"fontFamily"`
	TextMutedAligh string `json:"textAlign" bson:"textAlign"`
	Display        string `json:"display" bson:"display"`
}

type role struct {
	Fontsize      int    `json:"fontSize" bson:"fontSize"`
	TextTrasnform string `json:"textTransform" bson:"textTransform"`
	TextAligh     string `json:"textAlign" bson:"textAlign"`
	Display       string `json:"display" bson:"display"`
	LetterSpacing int    `json:"letterSpacing" bson:"letterSpacing"`
}

type photo struct {
	Width        int    `json:"width" bson:"width"`
	Height       int    `json:"height" bson:"height"`
	BorderRadius string `json:"borderRadius" bson:"borderRadius"`
	Margin       string `json:"margin" bson:"margin"`
}

type title struct {
	Fontsize      int    `json:"fontSize" bson:"fontSize"`
	TextTrasnform string `json:"textTransform" bson:"textTransform"`
	LetterSpacing int    `json:"letterSpacing" bson:"letterSpacing"`
	MarginBottom  int    `json:"marginBottom" bson:"marginBottom"`
}

type text struct {
	Fontsize   int    `json:"fontSize" bson:"fontSize"`
	LineHeight string `json:"lineHeight" bson:"lineHeight"`
	Color      string `json:"color" bson:"color"`
	Display    string `json:"display" bson:"display"`
}

type textMuted struct {
	Fontsize int    `json:"fontSize" bson:"fontSize"`
	Color    string `json:"color" bson:"color"`
}

type subTitle struct {
	Fontsize      int    `json:"fontSize" bson:"fontSize"`
	Color         string `json:"color" bson:"color"`
	Display       string `json:"display" bson:"display"`
	TextTrasnform string `json:"textTransform" bson:"textTransform"`
}

type link struct {
	Color      string `json:"color" bson:"color"`
	LineHeight int    `json:"lineHeight" bson:"lineHeight"`
}

// Style model for resume
type Style struct {
	Body      body      `json:"body" bosn:"body"`
	Display   display   `json:"display" bosn:"display"`
	Name      name      `json:"name" bosn:"name"`
	Role      role      `json:"role" bosn:"role"`
	Photo     photo     `json:"photo" bosn:"photo"`
	Title     title     `json:"title" bosn:"title"`
	Text      text      `json:"text" bosn:"text"`
	TextMuted textMuted `json:"textMuted" bosn:"textMuted"`
	SubTitle  subTitle  `json:"subTitle" bosn:"subTitle"`
	Link      link      `json:"link" bosn:"link"`
}
