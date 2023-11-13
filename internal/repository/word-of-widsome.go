package repository

import (
	"context"
	"math/rand"
)

type WordOfWidsomeRepo struct{}

func NewWordOfWidsomeRepo() *WordOfWidsomeRepo {
	return &WordOfWidsomeRepo{}
}

func (w WordOfWidsomeRepo) GetQuote(ctx context.Context) (string, error) { // error - for future implementation with some api or data storage
	return QuotesFromWordOfWidsome[rand.Intn(len(QuotesFromWordOfWidsome))], nil
}

var QuotesFromWordOfWidsome []string = []string{
	"Be more concerned with your character than your reputation, because your character is what you really are, while your reputation is merely what others think you are.",
	"To acquire knowledge, one must study; but to acquire wisdom, one must observe",
	"Talent is God given. Be humble. Fame is man-given. Be grateful. Conceit is self-given. Be careful.",
	"Don't give to anyone the power to put you down. Haters are losers pretending to be winners.",
	"Yesterday I was clever, so I wanted to change the world. Today I am wise, so I am changing myself.",
	"There are moments when troubles enter our lives and we can do nothing to avoid them. But they are there for a reason. Only when we have overcome them will we understand why they were there.",
	"A man only becomes wise when he begins to calculate the approximate depth of his ignorance.",
	"If you cannot find peace within yourself, you will never find it anywhere else.",
	"Knowing others is intelligence; knowing yourself is true wisdom. Mastering others is strength; mastering yourself is true power.",
	"A man must be big enough to admit his mistakes, smart enough to profit from them, and strong enough to correct them.",
	"Don’t depend too much on anyone in this world because even your own shadow leaves you when you are in darkness.",
	"Don't find fault, find a remedy.",
	"Kindness is the language which the deaf can hear and the blind can see.",
	"Experience is not what happens to you; it's what you do with what happens to you.",
	"We seem to gain wisdom more readily through our failures than through our successes. We always think of failure as the antithesis of success, but it isn't. Success often lies just the other side of failure.",
	"Always walk through life as if you have something new to learn, and you will.",
	"Smart men walked on the moon, daring men walked on the ocean floor, but wise men walk with God.",
	"I have held many things in my hands, and I have lost them all; but whatever I have placed in God's hands, that I still possess.",
	"The simple things are also the most extraordinary things, and only the wise can see them.",
	"The past is behind, learn from it. The future is ahead, prepare for it. The present is here, live it.",
	"There are two ways of spreading light: to be the candle or the mirror that reflects it.",
	"Everything is hard before it is easy.",
}
