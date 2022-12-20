package main

import (
    "fmt"
)

func main() {
    var testCases = []string{
        "abccccdd",
        "bananas",
        "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth",
    }
    for i := range testCases {
        fmt.Println(longestPalindrome(testCases[i]))
    }
}

func longestPalindrome(s string) int {
    m, c := make(map[byte]bool), 0 //мапа для поиска дублей
    for i := range s {
        if m[s[i]] { //если значение есть в мапе то инкризим каунтер и удаляем его из мапы
            c++
            delete(m, s[i]) //удаляем
        } else { //записываем тру если нечетное кол-во байт
            m[s[i]] = true
        }
    }
    if len(m) > 1 { //проверяем есть ли элементы в мапе, если остались, то можно вставить элемент в середину полиндрома
        return c*2 + 1
    }
    return c * 2 //возвращаем каунтер умноженный на 2
}
