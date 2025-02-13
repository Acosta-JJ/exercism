package blackjack

func ParseCard(card string) int {
    switch card {
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten":
        return 10
    case "jack":
        return 10
    case "queen":
        return 10
    case "king":
        return 10
    case "ace":
        return 11
    default:
        return 0
    }
}

func FirstTurn(card1, card2, dealerCard string) string {
    handScore := ParseCard(card1) + ParseCard(card2)
    dealerScore := ParseCard(dealerCard)

    switch {
    case handScore == 22:
        return "P"
    case handScore == 21 && dealerScore >= 10:
        return "S"
    case handScore == 21:
        return "W"
    case 17 <= handScore && handScore <= 20:
        return "S"
    case 12 <= handScore && handScore <= 16 && dealerScore >= 7:
        return "H"
    case 12 <= handScore && handScore <= 16:
        return "S"
    default:
        return "H"
    }
}

