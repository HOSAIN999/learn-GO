package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() { //در کد زیر میگوییم زمانی که برنامه اجرا شد پکیج کیبورد را باز کن و اگر خطایی بود ثبت کن و سپس برنامه را ببند
	err := keyboard.Open() //  این میگوید که بیا و کیبورد را باز کن و داخل err قرار بده(open یک خطا را بر میگرداند)
	if err != nil {        //اگر که err خطایی برگرداند که برابر با 0یا null نبود
		log.Fatal(err) //بیا و برنامه را متوقف کن(که بسیار نادر است)
	}
	defer func() { //برای این است که این فعل را الان اجرا نکند defer
		_ = keyboard.Close() //زمانی که خطایی رخ داد آن را نادیده بگیرد و پکیج کیبورد را ببندد
	}()
	coffees := make(map[int]string) //روش ساخت یک map
	coffees[1] = "Cappocino"
	coffees[2] = "coffee"
	coffees[3] = "Espresso"
	coffees[4] = "Latte"
	coffees[5] = "Mocha"
	coffees[6] = "Macchiato"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1-Cappocino")
	fmt.Println("2-coffee")
	fmt.Println("3-Espresso")
	fmt.Println("4-Latte")
	fmt.Println("5-Mocha")
	fmt.Println("6-Macchiato")
	fmt.Println("Q-Quite")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil { //این خط میاد و خطایی که پیش آمد رو به ما نشان میدهد
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q' {
			break
		}
		i, _ := strconv.Atoi(string(char))                    //در اینجا چون char یک نشانه است آن را به یک عدد صحیح تبدیل کردیم بعد به پایین میدهیم
		fmt.Println(fmt.Sprintf("You Choose %s", coffees[i])) //این خط کد مختصر کد های پایین است(در این جا کلا t را حذف کردیم)
		//t := fmt.Sprintf("You Choose %s", coffees[i]) //در این قسمت string در %s ذخیره میشود (اینجا گفتیم که برو و از داخل map اون گزینه رو بیار که int برابر با عدد وارد شده است)
		//t := fmt.Sprintf("You Choose %d",i)//در این قسمت int  در داخل %dذخیره میشود
		//t := fmt.Sprintf("You Choose %q",char)//(این قسمت میآید ویا همان key codeمورد نظر را به نشانه آن تبدیل میکند مثلا key code=49را به نشانه عدد 1در می آورد)
		//(%dداخل خود اعداد صحیح رو نگه میدارد)(%qداخل خود نشانه را نگه میدارد)

	}
	fmt.Println("Program Exist...")
}
