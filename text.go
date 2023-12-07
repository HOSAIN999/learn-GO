package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// func arrayfo() {

// 	var a [4]int
// 	fmt.Println("emp:", a)

// 	a[3] = 100
// 	fmt.Println("set:", a)
// 	fmt.Println("get:", a[3])

// 	fmt.Println("len:", len(a))

// 	b := [5]int{1, 2, 3, 4, 5}
// 	fmt.Println("dcl:", b)

// 	var twoD [2][3]int
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 3; j++ {
// 			twoD[i][j] = i + j
// 		}
// 	}
// 	fmt.Println("2d: ", twoD)

// }

// func helloIFELSE() {

// 	if 7%2 == 0 {
// 		fmt.Println("7 is even")
// 	} else {
// 		fmt.Println("7 is odd")
// 	}

// 	if 8%4 == 0 {
// 		fmt.Println("8 is divisible by 4")
// 	}

// 	if 7%2 == 0 || 8%2 == 0 {
// 		fmt.Println("either 8 or 7 are even")
// 	}

// 	if num := 9; num < 0 {
// 		fmt.Println(num, "is negative")
// 	} else if num < 10 {
// 		fmt.Println(num, "has 1 digit")
// 	} else {
// 		fmt.Println(num, "has multiple digits")
// 	}
// }

// func message() {

// 	var a = "initial"
// 	fmt.Println(a)

// 	var b, c int = 1, 2
// 	fmt.Println(b, c)

// 	var d = true
// 	fmt.Println(d)

// 	var e int
// 	fmt.Println(e)

// 	f := "apple"
// 	fmt.Println(f)
// }

// const s string = "constant"
// const h string = "hosain"

// func foring() {

// 	i := 1
// 	for i <= 3 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}

// 	for j := 7; j <= 9; j++ {
// 		fmt.Println(j)
// 	}

// 	for {
// 		fmt.Println("loop")
// 		break
// 	}

// 	for n := 0; n <= 5; n++ {
// 		if n%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(n)
// 	}
// }
// fmt.Println("hello world")
// fmt.Println("go" + "lang")
// fmt.Println("1+1 =", 23*10)
// fmt.Println("7.0/3.0 =", 78.0/10.0)
// fmt.Println(true && false)
// fmt.Println(true || false)
// fmt.Println(!true)
// fmt.Println(s)

// const n = 500000000

// const d = 3e20 / n
// fmt.Println(d)

// fmt.Println(int64(d))

// fmt.Println(math.Cos(n))
// fmt.Println(h + " rezaee")
// message()
// foring()

// for m := 0; m <= 10; m++ {
// 	if m%2 == 0 {
// 		continue
// 	}
// 	fmt.Println(m)
// }
// for {
// 	fmt.Println("loop")
// 	break

// }
// helloIFELSE()
// arrayfo()
// /////////////////////////////////////////شروع گلنگ با udemi////////////////////////////////////////////////////////////
// ما تمام کد های گلنگ را داخل یک packages قرار میدهیم

//  نکته دیگر هم این هست که از کتابخانه هایی استفاده می کند
//  import main
// 	"fmt"

//در زبان گلنگ اگر شما بخواهید که یک variableتعریف کنید به این شکل هست
// var a string
// const b = "hosain" //این برای این هست که بیرون از یک function باشد
// func main(){//در گلنگ به این شکل است که اولین func به نام main است و دیگر func در داخل این یکی صدا زده می شود
//  n := 20 این هم تعریف نوعی variable هست اما خلاصه تر
// }
//برای این که بتوانیم از یک فایل go دیگر یک func را صدا بزنیم به این روش عمل میکنیم
//1 ابتدا این دستور را در terminal ویژوال استودیو اجرا میکنیم go mod init myapp یک فایل به نام go.mod می سازد
//سپس در فایل main خودمان در import این کد را مینویسیم "myapp/نام فایل(doctor)"
//سپس هر جا که خواستیم یک func از فایل doctor را استفاده کنیم به این روش هست what to say := doctor.Intro()
//حالا برای این که بتوانیم چیزی که کاربر تایپ میکند را بخوانیم باید به این روش عمل کنیم
//reader := bufio.NewReader(os.Stdin) الان ما به وسیله این میتوانیم چیزهایی که کاربر تایپ میکند را بخوانیم البته که دو مورد کتابخانه هم اضافه شد
//به نام های "bufio","os"در قسمت import
//خوب حالا باید بگوییم که کاربر کجا تایپ کند و ما چه چیزی را از کاربر دریافت میکنیم و تا چه زمانی
//userInput, _ := reader.ReadString('\n')  در اینجا وقتی که \n را میزنیم یعنی کاربر دکمه enter را زده است و اطلاعات را برگردان
//fmt.Println(userInput)
//الان زمانی که ما در قسمت terminal دستور go run main.go را میزنیم مارا بیرون نمی اندازد و میتوانیم چیزی هم تایپ کنیم
//و بعد از تایپ و زدن enter چیزی را که تایپ کردیم را به ما برمیگرداند
//اما مشکلی که هست این که این فقط یک بار اتفاق می افتد وبعد از زدن enter مارا بیرون میاندازد حالا برای این که بتوانیم تا بی نهایت تایپ کنیم
//از یک حلقه for{} استفاده میکنیم و آن دو خط کد را داخل آن میگذاریم
// نحوه استفاده از شرط یا همان if else هم به این صورت است که
// if userInput == "quit" {
// 	break
// } else {
// 	fmt.Println(userInput)
// }
//به این وسیله یک عدد random میسازیم
// rand.Seed(time.Now().UnixNano())
/////////////////////////////////////تمرین گلنگ به وسیله پکیج doctor Udemy/////////////////////////////////////////
// reader := bufio.NewReader(os.Stdin)
// 	whatToSay := doctor.Intro()
// 	fmt.Println(whatToSay)
// 	for {
// 		fmt.Print("-> ")
// 		userInput, _ := reader.ReadString('\n')
// 		userInput = strings.Replace(userInput, "\r\n", "", -1) //این برای ویندوز
// 		userInput = strings.Replace(userInput, "\n", "", -1)   //این برای لینوکس و مک وfreeDSB
// 		if userInput == "quit" {
// 			break
// 		} else {
// 			fmt.Println(doctor.Response(userInput))
// 		}
// 	}
//////////////////////////////////////انواع روش های تعریف variable//////////////////////////////////////////
//  var firstnumber int
//  firstnumber = 2
//  var secondnumber = 3
//  thirdnumber := 4
////////////////////////////////////تمرین Udemy بازی حدس اعداد////////////////////////////////////////////////
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"time"
// )

// const prompt = "Press The Enter Ready"

// func main() {
// 	//بازی اعداد شما در این بازی به یک عدد فکر میکنید و بعد به شما مراحلی گفته می شود و در آخر ما عدد مورد نظر را حدس میزنیم
// 	rand.Seed(time.Now().UnixNano())به این وسیله ما یک عدد random به دست می آوریم
// 	firstnumber := rand.Intn(8) + 2 اعداد random بین 1 تا 10 است
// 	secondnumber := rand.Intn(8) + 2
// 	subtraction := rand.Intn(8) + 2
// 	answer := firstnumber*secondnumber - subtraction
// 	playTheGame(firstnumber, secondnumber, subtraction, answer) این جا هم function پایین را فراخوانی کردیم
// }

// func playTheGame(firstnumber, secondnumber, subtraction, answer int) { //ما اینجا یک function تعریف کردیم و به آن ورودی دادیم و عملیات ها را گفتیم انجام بده و در آخردر main آن را فراخوانی کردیم
// 	reader := bufio.NewReader(os.Stdin) //به این وسیله ما میتوانستیم چیزی که کاربر وارد کرده است را بخوانیم
// 	fmt.Println("Guess the Number Game")
// 	fmt.Println("---------------------")
// 	fmt.Println("")
// 	fmt.Println("think of a Number between 1 and 10", prompt) //یک عد بین 1 تا 10 انتخاب کنید
// 	reader.ReadString('\n') به این وسیله کاربر میتواند enter وارد کند و به گزینه بعدی برود
// 	fmt.Println("Muliply Your Number By", firstnumber, prompt) //عدد انتخابی را ضرب عدد مورد نظر کنید
// 	reader.ReadString('\n')
// 	fmt.Println("Now Multiply The Result By", secondnumber, prompt) //
// 	reader.ReadString('\n')
// 	fmt.Println("Divide The Result By The Number Your Originally thought Of", prompt)
// 	reader.ReadString('\n')
// 	fmt.Println("Now Substract", subtraction, prompt)
// 	reader.ReadString('\n')

// 	fmt.Println("The Answer is", answer) در آخر جواب رو چاپ کن
// }
///////////////////////////////////Scope یا همان سطح دسترسی//////////////////////////////////////////////////////
//فرض کنید که شما یک variables تعریف  میکنید اما برای سطح دسترسی به آن در زبان golang چیزهایی تعریف شده است مثال:
// var one string = "one"//زمانی که شما خارج از main تعریف کنیم در تمام جاها میتوانیم از آن استفاده کنیم

// func main() {

// }
//  func hello (){
// 	two := "two"//اما زمانی که داخل یک func تعریف کنیم فقط داخل همان در دسترس ما هست
// 	fmt.Println(two)
//  }
//نکته دیگر هم این است که شما مثلا یک package دیگر به نام doctor دارید و میخواهید که از variables & functions
//در packages main استفاده کنید اگر در داخل doctor شما variables & functions   حرف اول را با حروف بزرگ نام گذاری کرده باشید
//می توانید که به آنها دسترسی داشته باشید اما اگر  حرف اول را با حروف کوچک شروع کرده باشید حد دسترسی در package دیگر را ندارید
//:مثال
// package main

// import (
// 	"fmt"
// 	"myapp/doctor"
// )

// var one = "one" //چون خاج از function تعریف شده است در تمام function ها در دسترس است

//	func main() {
//		var two = "two" //چون داخل هست فقط داخل همین در دسترس هست
//		fmt.Println(two)
//		ScopePublice()
//	}
//
//	func ScopePublice() {
//		fmt.Println(one)
//		newString := doctor.Intro() //چون با حرف بزرگ شروع شده در این صفحه در دسترس هست Intro
//		fmt.Println("another page", newString)
//	}
/////////////////////////تمرین دوم ساخت variable & function در یک صفحه دیگر و فراخوانی در صفحه main وچاپ اطلاعات در یک خط/////////////
// var BlockVar = "skjvbs;jbvs;jrv;sorbv" //چون با حروف بزرگ هست در صفحه main قابل دسترس هست
// func PrintMe(s1, s2 string) {
// 	fmt.Println(s1, s2, BlockVar)
// }
//package main

// import (
// 	"fmt"
// 	"myapp/doctor"
// )

// var one = "one" //چون خاج از function تعریف شده است در تمام function ها در دسترس است

// func main() {
// 	var two = "two" //چون داخل هست فقط داخل همین در دسترس هست
// 	fmt.Println(two)
// 	ScopePublice()
// 	doctor.PrintMe(one, two)
// }
// func ScopePublice() {
// 	fmt.Println(one)
// 	newString := doctor.Intro() //چون با حرف بزرگ شروع شده در این صفحه در دسترس هست Intro
// 	fmt.Println("another page", newString)
// }
//////////تمرین کاربر چیزی تایپ میکند و پروژه آن را میگیرد دوباره چاپ میکند وبه کاربر نشان میدهد وبا نوشتن (quit)خارج میشود////////////
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// }
// reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print("-> ")
// 		userInput, _ := reader.ReadString('\n')                //به این وسیله وقتی که enter زده میشود می آید و چیزی که کاربر نوشته را می خواند
// 		userInput = strings.Replace(userInput, "\r\n", "", -1) //این جا ما می گوییم که چیزی که کاربر وارد کرده است را بخوان و enter که زده شد آن را پاک من و سپس یک خط هم پایین برو
// 		if userInput == "quit" {                               //اگر کاربر این کلمه را تایپ کرد از ترمینال بیرون بیا
// 			break
// 		} else { //در غیر این صورت هم چیزی که کاربر تایپ کرده را چاپ کن
// 			fmt.Println(userInput)
// 		}
// 	}
////////https://github.com/eiannone/keyboard//////////////////////go get -u github.com/eiannone/keyboard کد استفاده از دکمه های کیبورد
///////////////////////////////پروژه اعداد کلید//////////////////////////////////////////////////
// func main() { //در کد زیر میگوییم زمانی که برنامه اجرا شد پکیج کیبورد را باز کن و اگر خطایی بود ثبت کن و سپس برنامه را ببند
// 	err := keyboard.Open() //  این میگوید که بیا و کیبورد را باز کن و داخل err قرار بده(open یک خطا را بر میگرداند)
// 	if err != nil {        //اگر که err خطایی برگرداند که برابر با 0یا null نبود
// 		log.Fatal(err) //بیا و برنامه را متوقف کن(که بسیار نادر است)
// 	}
// 	defer func() { //برای این است که این فعل را الان اجرا نکند defer
// 		_ = keyboard.Close() //زمانی که خطایی رخ داد آن را نادیده بگیرد و پکیج کیبورد را ببندد
// 	}()

// 	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

// 	for {
// 		char, key, err := keyboard.GetSingleKey()
// 		if err != nil {//این خط میاد و خطایی که پیش آمد رو به ما نشان میدهد
// 			log.Fatal(err)
// 		}
// 		if key != 0 {
// 			fmt.Println("You Pressed", char, key)
// 		} else {

// 			fmt.Println("You Pressed", char)
// 		}
// 		if key == keyboard.KeyEsc {
// 			break
// 		}

// 	}
// 	fmt.Println("Program Exist...")

// }
/////////////////////////////////////////برای گرفتن یک ورودی از کاربر///////////////////////////////////////////////
// reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("what is your name: ")
// 	fmt.Print("-> ")
// 	userInput, _ := reader.ReadString('\n')
// 	userInput = strings.Replace(userInput,"\n", " ", -1) replace یعنی بیا کلمه ای که کاربر نوشته رو جایگزین جای خالی کن با زدن enter
// 	fmt.Println("your name is: ",userInput)
//////////////////////////////////////////پروژه منوی کافی شاپ //////////////////////////////////////////////////
// package main

// import (
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"github.com/eiannone/keyboard"
// )

// func main() { //در کد زیر میگوییم زمانی که برنامه اجرا شد پکیج کیبورد را باز کن و اگر خطایی بود ثبت کن و سپس برنامه را ببند
// 	err := keyboard.Open() //  این میگوید که بیا و کیبورد را باز کن و داخل err قرار بده(open یک خطا را بر میگرداند)
// 	if err != nil {        //اگر که err خطایی برگرداند که برابر با 0یا null نبود
// 		log.Fatal(err) //بیا و برنامه را متوقف کن(که بسیار نادر است)
// 	}
// 	defer func() { //برای این است که این فعل را الان اجرا نکند defer
// 		_ = keyboard.Close() //زمانی که خطایی رخ داد آن را نادیده بگیرد و پکیج کیبورد را ببندد
// 	}()
// 	coffees := make(map[int]string) //روش ساخت یک map
// 	coffees[1] = "Cappocino"
// 	coffees[2] = "coffee"
// 	coffees[3] = "Espresso"
// 	coffees[4] = "Latte"
// 	coffees[5] = "Mocha"
// 	coffees[6] = "Macchiato"

// 	fmt.Println("MENU")
// 	fmt.Println("----")
// 	fmt.Println("1-Cappocino")
// 	fmt.Println("2-coffee")
// 	fmt.Println("3-Espresso")
// 	fmt.Println("4-Latte")
// 	fmt.Println("5-Mocha")
// 	fmt.Println("6-Macchiato")
// 	fmt.Println("Q-Quite")

// 	for {
// 		char, _, err := keyboard.GetSingleKey()
// 		if err != nil { //این خط میاد و خطایی که پیش آمد رو به ما نشان میدهد
// 			log.Fatal(err)
// 		}
// 		if char == 'q' || char == 'Q' {
// 			break
// 		}
// 		i, _ := strconv.Atoi(string(char))                    //در اینجا چون char یک نشانه است آن را به یک عدد صحیح تبدیل کردیم بعد به پایین میدهیم
// 		fmt.Println(fmt.Sprintf("You Choose %s", coffees[i])) //این خط کد مختصر کد های پایین است(در این جا کلا t را حذف کردیم)
// 		//t := fmt.Sprintf("You Choose %s", coffees[i]) //در این قسمت string در %s ذخیره میشود (اینجا گفتیم که برو و از داخل map اون گزینه رو بیار که int برابر با عدد وارد شده است)
// 		//t := fmt.Sprintf("You Choose %d",i)//در این قسمت int  در داخل %dذخیره میشود
// 		//t := fmt.Sprintf("You Choose %q",char)//(این قسمت میآید ویا همان key codeمورد نظر را به نشانه آن تبدیل میکند مثلا key code=49را به نشانه عدد 1در می آورد)
// 		//(%dداخل خود اعداد صحیح رو نگه میدارد)(%qداخل خود نشانه را نگه میدارد)

// 	}
// 	fmt.Println("Program Exist...")
// }
////////////////////////////////////////برای گرفتن چند ورودی از کاربر و نشان دادن ورودی ها/////////////////////////////////////
// package main

// import (
// "bufio"
// "log"
// "fmt"
// "os"
// "strconv"
// "strings"
// "github.com/eiannone/keyboard"
// )

// var reader *bufio.Reader

// type User struct {اینجا ما یک struct تعریف کردیم که خصوصیات یک کاربر است
// 	UserName       string
// 	Age            int
// 	FavoriteNumber float64
// 	Ownadog        bool
// }

// func main() {
// 	var user User
// reader = bufio.NewReader(os.Stdin)این میاد و ورودی های ما رو میخونه
// 	user.UserName = readString("What is Your Name: ") //اینجا میاد سوال رو در قالب یک string به function ما میدهدوداخل username میگذارد

//	user.Age = readInt("How Old Are You?") //اینجا میاد سوال رو در قالب یک string به function ما میدهد
//	user.FavoriteNumber = readfloat("What is Your Favorite Number?")
//	user.Ownadog = readBool("Do You own a Dog (y/n)?")
//	fmt.Println(fmt.Sprintf("Your Name is %s and age is %d Years Old. Your favourite number is %2f.Own a Dog: %t.",
//		user.UserName,
//		user.Age,
//		user.FavoriteNumber,
//		user.Ownadog)) //اون %s یعنی از نوع string , %d یعنی از نوع عدد , %f یعنی برای اعداد اعشاری و اون عدد پشتش تعداد اعشاری هاش هست که قابل تغییر هست
//	//%t برای true & false  هست
//
// fmt.Print(*reader)
// }

// func prompt() {
// 	fmt.Print("-> ")
// }

// func readString(s string) string { //این خط میگه که یک string به عنوان ورودی میگیره و یک string به عنوان خروجی میدهد
// 	for {

// 		fmt.Println(s) //اینجا میگه که یک string به عنوان ورودی میگیرد و یک int به عنوان خروجی میدهد
// 		prompt()
// 		userInput, _ := reader.ReadString('\n')
// 		userInput = strings.Replace(userInput, "\r\n", "", -1)
// 		if userInput != "" { //اگر که کاربر اطلاعاتی وارد کرده است بیا و اون رو چاپ کن
// 			return userInput
// 		} else { // در غیر اینصورت بیا و این پیغام رو چاپ کن
// 			fmt.Println("Please enter a Value")
// 		}
// 	}
// }
// func readInt(s string) int { //اینجا میگه که یک string به عنوان ورودی میگیرد و یک int به عنوان خروجی میدهد
// 	for {
// 		fmt.Println(s) //اینجا اون string که به عنوان ورودی گرفته است را چاپ میکند که همان سوال است
// 		prompt()
// 		userInput, _ := reader.ReadString('\n') //اینجا اطلاعاتی که کاربر وارد میکند را میخواند
// 		userInput = strings.Replace(userInput, "\r\n", "", -1)

// 		num, err := strconv.Atoi(userInput) //این ورودی کاربر رو تبدیل میکند به int یا عدد
// 		if err != nil {                     //اگر خطا داشتیم بیا و این رو چاپ کن
// 			fmt.Println("Please enter a Whole Number")
// 		} else { //در غیر این صورت بیا و اون اطلاعاتی که از کاربر گرفتی و به عدد تبدیل کردی نشون بده
// 			return num
// 		}
// 	}
// }
// func readfloat(s string) float64 { //اینجا میگه که یک string به عنوان ورودی میگیرد و یک floate64 به عنوان خروجی میدهد
// 	for {
// 		fmt.Println(s) //اینجا اون string که به عنوان ورودی گرفته است را چاپ میکند که همان سوال است
// 		prompt()
// 		userInput, _ := reader.ReadString('\n')                //اینجا اطلاعاتی که کاربر وارد میکند را میخواند
// 		userInput = strings.Replace(userInput, "\r\n", "", -1) //در اینجا چیزی که کاربر وارد میکند را در خود ذخیره میکند تا بعدا استفاده کند

// 		num, err := strconv.ParseFloat(userInput, 64) //این ورودی کاربر رو تبدیل میکند به floate64 یا عدد
// 		if err != nil {                               //اگر خطا داشتیم بیا و این رو چاپ کن
// 			fmt.Println("Please enter a  Number")
// 		} else { //در غیر این صورت بیا و اون اطلاعاتی که از کاربر گرفتی و به عدد تبدیل کردی نشون بده
// 			return num
// 		}
// 	}
// }
// func readBool(s string) bool { //اینجا میگه که یک string به عنوان ورودی میگیرد و یک bool به عنوان خروجی میدهد
// 	err := keyboard.Open()این میاد و نشانه های کیبورد رو باز میکند
// 	if err != nil {
// 		log.Fatal()
// 	}
// 	defer func() {بی نام و نشان رو باید در آخرشون یک پرانتز باز و بسته بگذاریم function
// 		_ = keyboard.Close()
// 	}()
// 	for {
// 		fmt.Println(s) //اینجا اون string که به عنوان ورودی گرفته است را چاپ میکند که همان سوال است
// 		char, _, err := keyboard.GetSingleKey()
// 		if err != nil {
// 			log.Fatal()
// 		}اینجا اومدن و چند تا شرط گذاشتم برای این که وقتی رو کلمات y , n کلیک شد چه کار هایی انجام بشود
// 		if strings.ToLower(string(char)) == "y" && strings.ToLower(string(char)) == "n" {
// 			fmt.Println("Please type y or n")
// 		}else if char == 'n' || char == 'N'{
// 			return false
// 		}else if char == 'y' || char == 'Y'{
// 			return true
// 		}

// 	}
// }
////////////////////////////////////////////Basic Types(number(int, float,unit)string,bool)//////////////////////////////////////////////////
// var MyInt int
// var MyInt16 int16
// var MyInt32 int32
// var MyInt64 int64
// var MyUint uint (uint8,uint16,uint32,uint64)//برای اعداد صحیح استفاده میشود از 0 تا .... اما اعداد منفی را شامل نمی شود
// var MyFloat float (float32,float64)//برای اعداد اعشاری استفاده میشود
// var mystring string ="hosain"
// vat mybool bool (true,false)
/////////////////////////////////////////////aggregate type(array,struct)/////////////////////////////////////////
// func main(){

// 	var mystring [3]string//در زبان گولنگ به ندرت از آرایه استفاده میشود
// 	mystring[0] = "hosain"
// 	mystring[1] = "reza"
// 	mystring[2] = "hamed"
// 	fmt.Println("mystring",mystring[2])
// }
// type Car struct {
// 	NumberofTires int
// 	Luxery        bool
// 	BuketSeat     bool
// 	Model         string
// 	Make          string
// 	Year          int
// }

// func main(){
// var mycar car  //این یک روش نوشتن یک struct است
// mycar.NumberofTires = 4
// mycar.Luxery = false
// mycar.BuketSeat = true
// mycar.Model = "131Pride"
// mycar.Make = "iacnalc"
// mycar.Year = 2020
// mycar := Car {//این هم روش خلاصه تر هست که بیشتر استفاده میشود
//  NumberofTires : 4,
//  Luxery:  false,
//  BuketSeat:  true,
//  Model: "131Pride",
//  Make: "iacnalc",
//  Year: 2020,
// }
// fmt.Printf("My Car Is a %d %s %s",myCar.Year, myCar.Model, myCar.Make)
// }
//////////////////////////////////////////////reference types(pointers,slices,maps,functions,channels)////////////////////////////////////
//Pointers علامت های نشانه گذاری
// package main

// import "fmt"

// func main() {
// 	x := 10
// 	myfirstpointer := &x
// 	fmt.Println(x)
// 	fmt.Println(myfirstpointer)
// 	*myfirstpointer = 20
// 	fmt.Println(x)
// 	firstFunc(&x)
// 	fmt.Println(x)
// }
// func firstFunc(num *int) {
// 	*num = 30
// }
////////////////////////////////////////////////////////////////Slices////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	var animals []string
// 	animals = append(animals, "dog")
// 	animals = append(animals, "cat")
// 	animals = append(animals, "fish")
// 	animals = append(animals, "horse")
// 	fmt.Println(animals) //وقتی که یک slice رو اینطوری چاپ میکنیم به این صورت نشان میدهد[dog,cat,fish,horse]
// 	for _, x := range animals {
// 		fmt.Println(x) //وقتی به این صورت چاپ کنیم به این شکل نمایش میدهد
// 		//dog
// 	} //cat
// 	//fish
// 	//horse
// 	for i, x := range animals {
// 		fmt.Println(i, x) //وقتی به این صورت چاپ کنیم به این شکل نمایش میدهد
// 		//0 dog
// 	} //1 cat
// 	//2 fish
// 	//3 horse
// 	//اگر بخواهیم یک عنصر از slice را چاپ کنیم به این روش هست
// 	fmt.Println(animals[0]) //عدد قابل تغییر هست
// 	//اگر بخواهیم چند عنصر از یک slice را چاپ کنیم البته نه همه ی آن را به این روش هست
// 	fmt.Println(animals[0:2]) //اعداد قابل تغییر هست
// 	//اگر بخواهم ببینم چند عنصر در slice وجود دارد به این روش
// 	fmt.Println(len(animals))
// 	//ما میتوانیم slice ها را مرتب کنیم بر اساس حروف الفبا ابتدا میپرسیم که آیا مرتب شده است
// 	fmt.Println(sort.StringsAreSorted(animals)) //این یک جواب true, false برمیگرداند
// 	//حالا با این کد مرتب میکنیم
// 	sort.Strings(animals)
// 	fmt.Println(sort.StringsAreSorted(animals))
// 	fmt.Println(animals)
// 	//این ادامه حذف هست
// 	animals = deleteFormSlice(animals, 0) //این عدد مثل index میماند که چندمی را پاک کند
// 	fmt.Println(animals)
// }

// // برای حذف یک عنصر از slice باید به این شکل عمل کرد
// func deleteFormSlice(a []string, i int) []string {
// 	a[i] = a[len(a)-1]
// 	a[len(a)-1] = ""
// 	a = a[:len(a)-1]
// 	return a
// }
/////////////////////////////////////////ساختار map////////////////////////////////
//package main

// import "fmt"

// func main() {
// 	intmap := make(map[string]int) //نحوه ساخت یک map
// 	intmap["one"] = 1              //مقدار دهی به یک map
// 	intmap["two"] = 2
// 	intmap["three"] = 3
// 	intmap["four"] = 4
// 	intmap["five"] = 5
// 	for key, value := range intmap { //یک map را نمی توان مرتب کرد
// 		fmt.Println(key, value)
// 	}
// 	//delete(intmap, "two") //نحوه ی حذف یک عنصر از map

// 	for key, value := range intmap {
// 		fmt.Println(key, value)
// 	}
// 	el, ok := intmap["two"] //برای این که بدانیم که یک عنصر داخل map هست یا نه به این شکل عمل میکنیم
// 	if ok {
// 		fmt.Println(el, "is in the map")
// 	} else {
// 		fmt.Println(el, "is not in the map")
// 	}
// }
//////////////////////////////////////////////ساختار function///////////////////////////////////////////////
// package main

// import "fmt"
// func main(){
//    z:= addtwoNUMBER(20,30)
//    fmt.Println(z)
// }
// func addtwoNUMBER(X,Y int)(sum int){//این یک مدل function هست که در خود آن میگوییم که خروجی از چه نوع و به چه صورت هست
// 	sum=X+Y //البته این نوع خیلی کم استفاده می شود
// 	return
// }
// package main

// import "fmt"

// func main() {
// 	mytotal := sumMany(8, 35, 85, 47, 43, 25, 89, 157, 12)
// 	fmt.Println(mytotal)
// }
// func sumMany(nums ...int) int {//به این تابع Variadicچون معلوم نیست چند تا ورودی دارد
// 	total := 0
// 	for _, x := range nums {
// 		total = total + x
// 	}
// 	return total
// }
// package main
//  import (
// 	"fmt"
//  )
//  type animals struct {
// 	Name string
// 	Sound string
// 	numberOfLeg int
//  }
//  func(a *animals) Says(){//این یک تابع گیرنده است
// 	fmt.Printf("A %s says %s", a.Name, a.Sound)
// 	fmt.Println()
//  }
//  func(a *animals) howleg(){//این یک تابع گیرنده است
// 	fmt.Printf("A %s says %d", a.Name, a.numberOfLeg)
// 	fmt.Println()
//  }
//  func main(){
//   var dog animals
//   dog.Name = "dog"
//   dog.Sound ="woof"
//   dog.says()
//   cat := animals{
// 	Name: "cat",
// 	Sound:"meow",
// 	numberOfLeg:4,
//   }
//   cat.says()
//   cat.howleg()
//  }
////////////////////////////////////////////////channels /////////////////////////////////////////////////////
// package main

// import (
// 	"fmt"

// 	"github.com/eiannone/keyboard"
// )

// func main() {
// 	go doSomeThing("helloWorld")//این کلمه go یکی از فرق های زبان golang با دیگر زبان ها است به وسیله این کلمه شما میتوانید همزمان چندین کار را با هم انجام دهید
// 	fmt.Println("Another Message")//در زبان های دیگر مثل php شما یک کار را تمام میکنید و بعد کار بعدی و دوباره کار به اتمام رسید کار بعدی
// 	for {

//		}
//	}
//
//	func doSomeThing(s string) {
//		until := 0
//		for {
//			fmt.Println("s is", s)
//			until = until + 1
//			if until >= 5 {
//				break
//			}
//		}
//	}
// var KeyPressChan chan rune //نحوهی ساخت یک chanelle

// func main() {
// 	KeyPressChan = make(chan rune)//این قسمت برای ساخت یک channelle لازم است
// 	go lisenForKeyPress()
// 	fmt.Println("Press any Key,or q to quit")
// 	_ = keyboard.Open()
// 	defer func() {
// 		keyboard.Close()
// 	}()
// 	for {
// 		char, _, _ := keyboard.GetSingleKey()
// 		if char == 'q' || char == 'Q' {
// 			break
// 		}
// 		KeyPressChan <- char
// 	}
// }
// func lisenForKeyPress() {
// 	for {
// 		key := <-KeyPressChan
// 		fmt.Println("You Pressed", string(key))
// 	}
// }
/////////////////////////////////////////]Interface////////////////////////////////////////////////////
// package main
// type animals struct {
//   says()
//   HowManyLeg()
// }
// type dog struct{
//   Name string
//   Sound string
//   numberOfLeg int
// }
// func (d *dog) says() string{
//   return d.Sound
// }
// func (d *dog) HowManyLegs() int{
//   return c.numberOfLeg
// }
// type cat struct{
//   Name string
//   Sound string
//   numberOfLeg int
//   Hastail bool
// }
// func (c *cat) HowManyLegs() int{
//   return c.numberOfLeg
// }
// func (c *cat) says() string{
//   return d.Sound
// }
// func main() {
//    Dog := dog{
//     Name: "dog",
//     Sound: "Woof",
//     numberOfLeg: 4,
//    }
//    Riddle(&Dog)
//    Cat := cat{
//     Name: "cat",
//     Sound: "meow",
//     numberOfLeg: 4,
//     Hastail: true,
//    }
//    Riddle(&Cat)
// }
// func Riddle(a animals){
//    riddle := fmt.Sprintf(`this is animal says "%s" and has "%d" leg. What animal is it`, a.says(), a.HowManyLegs())
//    fmt.Println(riddle)
// }
///////////////////////////////////////////////Composition////////////////////////////////////////////////////////////
// type Vehicle struct {
// 	numberOfWheels    int
// 	NumberOfPassenger int
// }
// type Car struct {
// 	Make       string
// 	Model      string
// 	Year       int
// 	IsElectric bool
// 	IsHybrid   bool
// 	Vehicle    Vehicle
// }

// func (v Vehicle) showDetails() {
// 	fmt.Println("Number of Passenger", v.NumberOfPassenger)
// 	fmt.Println("Number of Wheels", v.numberOfWheels)
// }
// func (c Car) show() {
// 	fmt.Println("Make:", c.Make)
// 	fmt.Println("Model:", c.Model)
// 	fmt.Println("Year:", c.Year)
// 	fmt.Println("IsElectric:", c.IsElectric)
// 	fmt.Println("IsHybrid:", c.IsHybrid)
// 	c.Vehicle.showDetails()
// }
// func main() {
// 	suv := Vehicle{
// 		numberOfWheels:    4,
// 		NumberOfPassenger: 6,
// 	}
// 	volvoXC90 := Car{
// 		Make:       "Volvo",
// 		Model:      "XC90 T8",
// 		Year:       2021,
// 		IsElectric: false,
// 		IsHybrid:   true,
// 		Vehicle:    suv,
// 	}
// 	volvoXC90.show()
// 	fmt.Println()
// 	teslaModelX := Car{
// 		Make:       "Tesla",
// 		Model:      "Model X",
// 		Year:       2022,
// 		IsElectric: true,
// 		IsHybrid:   true,
// 		Vehicle:    suv,
// 	}
// 	teslaModelX.Vehicle.NumberOfPassenger = 7
// 	teslaModelX.Vehicle.numberOfWheels = 6
// 	teslaModelX.show()
// }
/////////////Exported vs Unexported قابلیت دسترسی به وسیله استفاده از حروف بزرگ در ابتدای variable & function & type & ....////////////////
// package staff//این یک فایل جدا گانه است این دو فایل با هم در ارتباط هستند

// var OverpaidLimit = 5000
// var UnderpaidLimit = 2000

// type Employee struct {//این همانند کلاس است که از چند خصوصیت تشکیل میشود
// 	FirstName string
// 	LastName  string
// 	Salary    int
// 	FullTime  bool
// }
// type Office struct {//این یک نوعی است که از جنس بالایی است
// 	AllStaff []Employee
// }

// func (e *Office) All() []Employee {//این یک function از نوع office هست که خروجی آن از نوع Employee است
// 	return e.AllStaff
// }
// func (e *Office) OverPaid() []Employee {//این یک function از نوع office هست که خروجی آن از نوع Employee است
// 	var overpaid []Employee
// 	for _, x := range e.AllStaff {
// 		if x.Salary > OverpaidLimit {
// 			overpaid = append(overpaid, x)
// 		}
// 	}
// 	return overpaid
// }
// func (e *Office) UnderPaid() []Employee {//این یک function از نوع office هست که خروجی آن از نوع Employee است
// 	var underpaid []Employee
// 	for _, x := range e.AllStaff {
// 		if x.Salary > UnderpaidLimit {
// 			underpaid = append(underpaid, x)
// 		}
// 	}
// 	return underpaid
// }
// package main//این دو فایل با هم در ارتباط هستند

// import (
// 	"log"
// 	"myapp/staff"
// )

// var employee = []staff.Employee{//ما به این روش هم می توانیم یک نوع را مقدار دهی کنیم البته این نوع در یک فایل دیگر هست و ما در این فایل مقدار داده ایم
// 	{FirstName: "Hosain", LastName: "Rezaee", Salary: 12000, FullTime: true},
// 	{FirstName: "Hamed", LastName: "Rezaee", Salary: 10000, FullTime: true},
// 	{FirstName: "Vahid", LastName: "Rezaee", Salary: 8000, FullTime: true},
// 	{FirstName: "Zahra", LastName: "Rezaee", Salary: 6000, FullTime: true},
// 	{FirstName: "Hesam", LastName: "Rezaee", Salary: 4000, FullTime: false},
// }

// func main() {
// 	myStaff := staff.Office{
// 		AllStaff: employee,
// 	}
// 	log.Println(myStaff.All())
// 	staff.OverpaidLimit = 7000//به این روش هم میتوانیم variable که در یک صفحه دیگر مقدار دارد را مقدار جدید بدهیم
// 	staff.UnderpaidLimit = 1000
// 	log.Println(myStaff.OverPaid())
// 	log.Println(myStaff.UnderPaid())

// }
///////////////////////////////////////////for انواع حلقه/////////////////////////////////////////////////////////////////
// func main(){

// 	for i := 0; i >= 10; i++ {
// 		fmt.Println("i is: ",i)
// 	}
// 		rand.Seed(time.Now().UnixNano())//این یک عدد random میسازد
// 		i := 1000
// 		for i > 100 {
// 			i = rand.Intn(1000) + 1//این جا میگوید که اعداد random که از1000بیشتر نباشد را با یک جمع کن
// 			fmt.Println("i is", i)
// 			if i > 100 {//درایجا هم می گوید که اگر بزرگتر از 100 بود ادامه بده اما اگر کوچکتر بود خارج شو
// 				fmt.Println("i is", i, "so loop going")
// 			}
// 		}
// 		fmt.Println("Got", i, "broken")
// }
// package logger//این دو فایل با هم هستند

// import "log"

// func LisenForLog(ch chan string) {
// 	for {//یک حلقه بینهایت یا infinite Loop
// 		mgs := <-ch
// 		log.Print(mgs)
// 	}
// }
// package main//این دو فایل با هم هستند

// import (
// 	"bufio"
// 	"fmt"
// 	"myapp/logger"
// 	"os"
// 	"time"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	ch := make(chan string)

// 	go logger.LisenForLog(ch)

// 	fmt.Println("EnterSomething")
// 	for i := 0; i < 5; i++ {
// 		fmt.Print("-> ")
// 		Input, _ := reader.ReadString('\n')
// 		ch <- Input
// 		time.Sleep(time.Second)
// 	}
// 	for i := 1; i <= 10; i++ {//این یک حلقه تو در تو استNestedLoop
// 		fmt.Println("i is", i, )
// 		for j := 1; j <= 3; j++ {
// 			fmt.Println("j:", j, )
// 		}
// 		fmt.Println()
// 	}
// }
////////////////////////////////////////////////menu coffeeshop به روش دیگر/////////////////////////////////////////
//  package main

//  import (
//  	"fmt"
//  	"log"
//  	"strconv"

//  	"github.com/eiannone/keyboard"
//  )

//  func main() { //در کد زیر میگوییم زمانی که برنامه اجرا شد پکیج کیبورد را باز کن و اگر خطایی بود ثبت کن و سپس برنامه را ببند
//  	err := keyboard.Open() //  این میگوید که بیا و کیبورد را باز کن و داخل err قرار بده(open یک خطا را بر میگرداند)
//  	if err != nil {        //اگر که err خطایی برگرداند که برابر با 0یا null نبود
//  		log.Fatal(err) //بیا و برنامه را متوقف کن(که بسیار نادر است)
//  	}
//  	defer func() { //برای این است که این فعل را الان اجرا نکند defer
//  		_ = keyboard.Close() //زمانی که خطایی رخ داد آن را نادیده بگیرد و پکیج کیبورد را ببندد
//  	}()
//  	coffees := make(map[int]string) //روش ساخت یک map
//  	coffees[1] = "Cappocino"
//  	coffees[2] = "coffee"
//  	coffees[3] = "Espresso"
//  	coffees[4] = "Latte"
//  	coffees[5] = "Mocha"
//  	coffees[6] = "Macchiato"

//  	fmt.Println("MENU")
//  	fmt.Println("----")
//  	fmt.Println("1-Cappocino")
//  	fmt.Println("2-coffee")
//  	fmt.Println("3-Espresso")
//  	fmt.Println("4-Latte")
//  	fmt.Println("5-Mocha")
//  	fmt.Println("6-Macchiato")
//  	fmt.Println("Q-Quite")
//     char := ' '
//  	for char != 'q' && char != 'Q'{//اینجا گفتیم که اگر این دکمه ها نبود بیا وکار های پایین رو انجام بده
//  		char, _, err = keyboard.GetSingleKey()
//  		if err != nil { //این خط میاد و خطایی که پیش آمد رو به ما نشان میدهد
//  			log.Fatal(err)
//  		}

//  		i, _ := strconv.Atoi(string(char))
// 		_, ok := coffees[i]  //اینجا میایم و یک ok میسازیم برابر با coffees و یک شرط میگذاریم
// 		if ok {
// 			fmt.Println(fmt.Sprintf("You Choose %s", coffees[i])) //این خط کد مختصر کد های پایین است(در این جا کلا t را حذف کردیم)
// 		}

// 		if _, ok := coffees[i]; ok {//این همان کاری را میکند که کد بالا انجام میدهدولی مختصر است
// 			fmt.Println(fmt.Sprintf("You Choose %s", coffees[i])) //این خط کد مختصر کد های پایین است(در این جا کلا t را حذف کردیم)
// 		}          //در اینجا چون char یک نشانه است آن را به یک عدد صحیح تبدیل کردیم بعد به پایین میدهیم
//  		//t := fmt.Sprintf("You Choose %s", coffees[i]) //در این قسمت string در %s ذخیره میشود (اینجا گفتیم که برو و از داخل map اون گزینه رو بیار که int برابر با عدد وارد شده است)
//  		//t := fmt.Sprintf("You Choose %d",i)//در این قسمت int  در داخل %dذخیره میشود
//  		//t := fmt.Sprintf("You Choose %q",char)//(این قسمت میآید ویا همان key codeمورد نظر را به نشانه آن تبدیل میکند مثلا key code=49را به نشانه عدد 1در می آورد)
//  		//(%dداخل خود اعداد صحیح رو نگه میدارد)(%qداخل خود نشانه را نگه میدارد)

//  	}
//  	fmt.Println("Program Exist...")
//  }
// for i := 0; i <= 100; i++ {//این کد میاد و 101 بار انجام میشود تا حلقه به پایان برسد
// 	if i % 7==0 {
// 		fmt.Println(i)
// 	}
// }
// for i := 0; i <= 100; i = i + 7 {//این کد هم کار بالایی را انجام میدهد ولی فقط 15 بار انجام میشودو از نظر سرعت خیلی بهتر است
// 	fmt.Println(i)
// }
// package main

// import (
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"github.com/eiannone/keyboard"
// )

// func main() { //در کد زیر میگوییم زمانی که برنامه اجرا شد پکیج کیبورد را باز کن و اگر خطایی بود ثبت کن و سپس برنامه را ببند
// 	err := keyboard.Open() //  این میگوید که بیا و کیبورد را باز کن و داخل err قرار بده(open یک خطا را بر میگرداند)
// 	if err != nil {        //اگر که err خطایی برگرداند که برابر با 0یا null نبود
// 		log.Fatal(err) //بیا و برنامه را متوقف کن(که بسیار نادر است)
// 	}
// 	defer func() { //برای این است که این فعل را الان اجرا نکند defer
// 		_ = keyboard.Close() //زمانی که خطایی رخ داد آن را نادیده بگیرد و پکیج کیبورد را ببندد
// 	}()
// 	coffees := make(map[int]string) //روش ساخت یک map
// 	coffees[1] = "Cappocino"
// 	coffees[2] = "coffee"
// 	coffees[3] = "Espresso"
// 	coffees[4] = "Latte"
// 	coffees[5] = "Mocha"
// 	coffees[6] = "Macchiato"

// 	fmt.Println("MENU")
// 	fmt.Println("----")
// 	fmt.Println("1-Cappocino")
// 	fmt.Println("2-coffee")
// 	fmt.Println("3-Espresso")
// 	fmt.Println("4-Latte")
// 	fmt.Println("5-Mocha")
// 	fmt.Println("6-Macchiato")
// 	fmt.Println("Q-Quite")

// 	for ok := true; ok; ok = char != 'q' || char != 'Q'{//این هم روش دیگری برای از حلقه هست ولی شبیه به do while که کمی خواندن آن سخت تر است
// 		char, _, err := keyboard.GetSingleKey()
// 		if err != nil { //این خط میاد و خطایی که پیش آمد رو به ما نشان میدهد
// 			log.Fatal(err)
// 		}
//
// 		i, _ := strconv.Atoi(string(char))                    //در اینجا چون char یک نشانه است آن را به یک عدد صحیح تبدیل کردیم بعد به پایین میدهیم
// 		fmt.Println(fmt.Sprintf("You Choose %s", coffees[i])) //این خط کد مختصر کد های پایین است(در این جا کلا t را حذف کردیم)
// 		//t := fmt.Sprintf("You Choose %s", coffees[i]) //در این قسمت string در %s ذخیره میشود (اینجا گفتیم که برو و از داخل map اون گزینه رو بیار که int برابر با عدد وارد شده است)
// 		//t := fmt.Sprintf("You Choose %d",i)//در این قسمت int  در داخل %dذخیره میشود
// 		//t := fmt.Sprintf("You Choose %q",char)//(این قسمت میآید ویا همان key codeمورد نظر را به نشانه آن تبدیل میکند مثلا key code=49را به نشانه عدد 1در می آورد)
// 		//(%dداخل خود اعداد صحیح رو نگه میدارد)(%qداخل خود نشانه را نگه میدارد)

// 	}
// 	fmt.Println("Program Exist...")
// }
///////////////////////////////////////////////////////بازی سنگ.کاغذ.قیچی روش اول////////////////////////////////////////////////////
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"os/exec"
// 	"runtime"
// 	"strings"
// 	"time"
// )

// const ( //سه تا ثابت داریم برای سنگ و کاغذ وقیچی
// 	ROCK     = 0
// 	PAPER    = 1
// 	SCISSORS = 2
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano()) //اینجا یک عدد صحیح به صورت random میسازیم
// 	playerChoice := ""
// 	playerValue := -1  //مقدار اولیه بازیکن باید 1- باشد
// 	playerScore := 0   //تعداد پیروزی های بازیکن
// 	computerScore := 0 //تعداد پیروزی های کامپیوتر
// 	reader := bufio.NewReader(os.Stdin)
// 	clearScreen() //به این وسیله ترمینال را پاک میکنیم
// 	fmt.Println("ROCK,PAPER & SCISSORS")
// 	fmt.Println("---------------------")
// 	fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
// 	fmt.Println("")
// 	for i := 1; i <= 3; i++ { //این حلقه را میزنیم تا بازی سه بار ادامه پیدا کند
// 		fmt.Println()
// 		fmt.Println("Round", i) //تعداد دفعاتی که بازی انجام شده است
// 		fmt.Println("-------")
// 		fmt.Print("Please enter rock,paper or scissors -> ")
// 		playerChoice, _ = reader.ReadString('\n')                    //اون چیزی که کاربر تایپ کرده را میخوانیم
// 		playerChoice = strings.Replace(playerChoice, "\r\n", "", -1) //اون چیزی که کاربر نوشته را مینویسیم
// 		computerValue := rand.Intn(3)                                //اینجا میگیم که اعداد بین یک تا سه باشد
// 		if playerChoice == "rock" {                                  //اینجا شرط میگذاریم که اگر کاربر تایپ کرده بود سنگ یا کاغذ و یا قیچی مساوی با چه مقدار ثابت باشه
// 			playerValue = ROCK
// 		} else if playerChoice == "paper" {
// 			playerValue = PAPER
// 		} else if playerChoice == "scissors" {
// 			playerValue = SCISSORS
// 		} else {
// 			playerValue = -1
// 		}
// 		fmt.Println()
// 		fmt.Println("Player chose", strings.ToUpper(playerChoice)) //اون چیزی که کاربر تایپ کرده رو به صورت حروف بزرگ تبدیل میکنیم

// 		switch computerValue { //این جا هم مقادیری که کامپیوتر به صورت random میگیره رو تبدیل میکنیم به سنگ کاغذ و قیچی
// 		case ROCK:
// 			fmt.Println("Computer Choose ROCK")
// 			break
// 		case PAPER:
// 			fmt.Println("Computer Choose PAPER")
// 			break
// 		case SCISSORS:
// 			fmt.Println("Computer Choose SCISSORS")
// 			break
// 		default:
// 		}
// 		fmt.Println()
// 		if playerValue == computerValue { //اینجا اگر مقدار کاربر با مقدار بازیکن برابر بود مساوی اعلام میکنه و بعد هم یک دور کم میکند
// 			fmt.Println("It is Draw")
// 			i--
// 		} else {
// 			switch playerValue { //اینجا مقدار های کاربر و کامپیوتر مقایسه میشوند و برنده هر دور اعلام میشود
// 			case ROCK:
// 				if computerValue == PAPER {
// 					computerScore = ComputerWins(computerScore)
// 				} else {
// 					playerScore = PlayerWins(playerScore)
// 				}
// 				break
// 			case PAPER:
// 				if computerValue == SCISSORS {
// 					computerScore = ComputerWins(computerScore)
// 				} else {
// 					playerScore = PlayerWins(playerScore)
// 				}
// 				break
// 			case SCISSORS:
// 				if computerValue == ROCK {
// 					computerScore = ComputerWins(computerScore)
// 				} else {
// 					playerScore = PlayerWins(playerScore)
// 				}
// 				break
// 			default:
// 				fmt.Println("Invalide Choise") //اگر مقداری که کاربر وارد کرده غیر از اون مقدار های ثابت باشه خطا میدهد و یک دور کم میکند
// 				i--
// 			}
// 		}
// 	}
// 	fmt.Println("Final Score")
// 	fmt.Println("-----------")
// 	fmt.Printf("Player %d/3, Computer %d/3", playerScore, computerScore) //اینجا هم تعداد برد بازی کن وکامپیوتر را نشان میدهد
// 	fmt.Println()
// 	if playerScore > computerScore { //اینجا هم بردهای بازی کن و کامپیوتر را مقایسه میکند و برنده را اعلام میکند
// 		fmt.Println("Player Wins Game!")
// 	} else {
// 		fmt.Println("Computer Wins Game!")

// 	}
// }
// func clearScreen() { //این همون فعل پاک کردن ترمینال هست
// 	if strings.Contains(runtime.GOOS, "windows") {
// 		// windows
// 		cmd := exec.Command("cmd", "/c", "cls")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	} else {
// 		// linux or mac
// 		cmd := exec.Command("clear")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	}
// }
// func ComputerWins(score int) int { //این هم تعداد پیروزی های کامپیوتر را در خود نگه میدارد و در پایان هر دور نشان میدهد
// 	fmt.Println("Computer wins!")
// 	return score + 1
// }
// func PlayerWins(score int) int { //این هم تعداد پیروزی های بازی کن را در خود نگه میدارد و در پایان هر دور نشان میدهد
// 	fmt.Println("Player wins!")
// 	return score + 1
// }
//////////////////////////////////////////////////////نحوه استفاده از chan به همراه select////////////////////////////////////////////////////
// package main

// import (
// 	"fmt"
// 	"time"
// )

// var chan1 = make(chan string)
// var chan2 = make(chan string)

// func task1() {
// 	time.Sleep(1 * time.Second)
// 	chan1 <- "One"
// }

// func task2() {
// 	time.Sleep(2 * time.Second)
// 	chan2 <- "Two"
// }

// func main() {
// 	task1()
// 	task2()
// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-chan1:
// 			fmt.Println("Received", msg1)
// 		case msg2 := <-chan2:
// 			fmt.Println("Received", msg2)
// 		}
// 	}
// }
//////////////////////////////////////////////////////بازی سنگ.کاغذ.قیچی روش دوم به وسیله Select ,chan///////////////////////////////////////////////////
// package main

// import "myapp/game"

// func main() {
// 	displaychan := make(chan string)
// 	roundchan := make(chan int)
// 	game := game.Game{
// 		DisplayChan: displaychan,
// 		RoundChan:   roundchan,
// 		Round: game.Round{
// 			RoundNumber:   0,
// 			PlayScore:     0,
// 			ComputerScore: 0,
// 		},
// 	}
// 	go game.Rounds()   //این تعداد راند های بازی رو عوض میکند
// 	game.ClearScreen() //این که ترمینال رو پاک میکند
// 	game.PrintIntro()  //اطلاعات اولیه بازی را چاپ میکند
// 	for {              //این جا هم یکی یکی به دور های بازی اضافه میشود تا به سه دور برسد و تمام شود
// 		game.RoundChan <- 1
// 		<-game.RoundChan
// 		if game.Round.RoundNumber > 3 {
// 			break
// 		}
// 		if !game.PlayRound() { //این شرط میگه اگر که یک دور به هر دلیلی انجام نشد یکی کم کن و این دور رو دوباره بازی کنند
// 			game.RoundChan <- -1
// 			<-game.RoundChan
// 		}
// 	}
// 	game.PrintSummery() //این اطلاعات بازی را در آخر نشان میدهد
// }
// //این دو فایل با هم هستند
// package game

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"os/exec"
// 	"runtime"
// 	"strings"
// 	"time"
// )

// const (
// 	ROCK     = 0
// 	PAPER    = 1
// 	SCISSORS = 2
// )

// type Game struct {
// 	DisplayChan chan string
// 	RoundChan   chan int
// 	Round       Round
// }
// type Round struct {
// 	RoundNumber   int
// 	PlayScore     int
// 	ComputerScore int
// }

// var reader = bufio.NewReader(os.Stdin)

// func (g *Game) Rounds() {
// 	for {
// 		select {
// 		case round := <-g.RoundChan:
// 			g.Round.RoundNumber = g.Round.RoundNumber + round
// 			g.RoundChan <- 1
// 		case msg := <-g.DisplayChan:
// 			fmt.Println(msg)
//      g.DisplayChan <- ""
// 		}
// 	}
// }
// func (g *Game) ClearScreen() { //این همون فعل پاک کردن ترمینال هست
// 	if strings.Contains(runtime.GOOS, "windows") {
// 		// windows
// 		cmd := exec.Command("cmd", "/c", "cls")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	} else {
// 		// linux or mac
// 		cmd := exec.Command("clear")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	}
// }
// func (g *Game) PrintIntro() {
//  g.DisplayChan <- fmt.Sprintf(`
//Rock,Paper & Scissors
//---------------------
//Game is Played for three rounds, and best of threewins the game.Good luck!
//`)
//   <-g.DisplayChan
// }
// func (g *Game) PlayRound() bool { //این فعل انجام بازی هست که به ما درست یا غلط برمیگرداند
// 	rand.Seed(time.Now().UnixNano())
// 	playerValue := -1 //این مقدار منفی یک هست چون مقادیر ثابت ما از صفر شروع میشود
// g.DisplayChan <- fmt.Sprintf(`
// Round %d
// -------
// `,g.Round.RoundNumber)
// <- g.DisplayChan

// 	fmt.Print("Please Enter Rock, Paper & Scissors-> ")
// 	playerChoice, _ := reader.ReadString('\n')
// 	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
// 	computerValue := rand.Intn(3)

// 	if playerChoice == "rock" {
// 		playerValue = ROCK
// 	} else if playerChoice == "paper" {
// 		playerValue = PAPER
// 	} else if playerChoice == "scissors" {
// 		playerValue = SCISSORS
// 	}
//  g.DisplayChan <- ""
// 	<-g.DisplayChan
// 	g.DisplayChan <- fmt.Sprintf("Player Chose %s", strings.ToUpper(playerChoice))
//  <-g.DisplayChan
// 	switch computerValue {
// 	case ROCK:
//  g.DisplayChan <- "Computer Chose ROCK"
// 	<-g.DisplayChan
// 		break
// 	case PAPER:
//  g.DisplayChan <- "Computer Chose PAPER"
// 	<-g.DisplayChan
// 		break
// 	case SCISSORS:
// 	  g.DisplayChan <- "Computer Chose SCISSORS"
// 	<-g.DisplayChan
// 		break
// 	default:
// 	}
// g.DisplayChan <-""
// <-g.DisplayChan
// 	if playerValue == computerValue {
// 		g.DisplayChan <- "It is a draw"
//      <-g.DisplayChan
// 		return false
// 	} else {
// 		switch playerValue {
// 		case ROCK:
// 			if computerValue == PAPER {
// 				g.computerWins()
// 			} else {
// 				g.playerWins()
// 			}
// 			break
// 		case PAPER:
// 			if computerValue == SCISSORS {
// 				g.computerWins()
// 			} else {
// 				g.playerWins()
// 			}
// 			break
// 		case SCISSORS:
// 			if computerValue == ROCK {
// 				g.computerWins()
// 			} else {
// 				g.playerWins()
// 			}
// 			break
// 		default:
// 			g.DisplayChan <- "Invalid Choice"
//          <-g.DisplayChan
// 			return false
// 		}
// 	}
// 	return true
// }
// func (g *Game) computerWins() { //این فعل می آید و یکی به تعداد برد های کامپیوتر اضافه میکند
// 	g.Round.ComputerScore++
// 	g.DisplayChan <- "Computer wins!"
//  <-g.DisplayChan
// }
// func (g *Game) playerWins() { //این یکی به تعداد بردهای کاربر اضافه میکند
// 	g.Round.PlayScore++
// 	g.DisplayChan <- "Player wins!"
//  <-g.DisplayChan
// }
// func (g *Game) PrintSummery() {
// g.DisplayChan <-fmt.Sprintf(`
// Final Score
// -----------
// Player: %d/3, Computer %d/3,
// `, g.Round.PlayScore, g.Round.ComputerScore)
// <-g.DisplayChan
// 	if g.Round.PlayScore > g.Round.ComputerScore {
//  g.DisplayChan <- "Player Win Game!"
// 	<-g.DisplayChan
// 	} else {
// 	  g.DisplayChan <-Computer Win Game!"
//    <-g.DisplayChan
// 	}
// }
///////////////////////////////////////////ریاضی در زبان گولنگ چگونه است///////////////////////////////////////////////
// در ابتدا این که محاسبات ریاضی به ترتیب اولویت انجام می شود یعنی ابتداداخل پرانتز ,بعد ضرب, بعد تقسیم ,بعدجمع وبعد تفریق
//answer := 3 + 4 * 5 = 23 چون ابتدا ضرب انجام میشود و بعد جمع
// اما اگر از پرانتز استفاده کنیم ابتدا عملیات های داخل پرانتز انجام میشود
//answer := (3 + 4) * 5 = 35 //چون اول داخل پرانتز انجام می شود
//برای این که اعداد دقیق تری بدست بیارید سعی کنید اعداد را به صورت اعشاری بنویسید
// var Radius = 12.0
//area := math.pi * radius * radius (pi نشانه عدد پی هست همان عدد 3/14)
//half := 1.0 / 2.0 این باعث میشود اگر عددی اعشاری شد نشان بدهد و اگر نگذاریم آن را صفر میکند و یا اعشار را گرد میکند
//power := math.Pow(3.0, 2.0) //pow برای به توان رساندن است
//remainder := 50 % 3 علامت % برای نشان دادن باقیمانده است که معمولا در تقسیم استفاده میشود
//اما علامت % کاربرد هایی خوبی در کوتاه کردن کد هم دارند که یک مثال در زیر زده شده است
// سعی کنید از پرانتز ها نیز در کد های خو استفاده کنید زیرا خوانا تر میشود و در آینده اگر بخواهید استفاده کنید بهتر است
// second := 31
// minutes := 1
// if (minutes <59)&&((second +1) > 59) {
// 	minutes++
// }
// fmt.Println(minutes)
///////////////////////////////////////////////////////بازی سنگ.کاغذ.قیچی روش سوم با استفاده از عملگر % باقی مانده////////////////////////////////////////////////////
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"os/exec"
// 	"runtime"
// 	"strings"
// 	"time"
// )

// const ( //سه تا ثابت داریم برای سنگ و کاغذ وقیچی
// 	ROCK     = 0 // beats scissors.(scissors + 1) % 3 = 0
// 	PAPER    = 1 //  beats Rock.(Rock + 1) % 3 = 1
// 	SCISSORS = 2 //  beats Paper.(Paper + 1) % 3 = 2
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano()) //اینجا یک عدد صحیح به صورت random میسازیم
// 	playerChoice := ""
// 	playerValue := -1  //مقدار اولیه بازیکن باید 1- باشد
// 	playerScore := 0   //تعداد پیروزی های بازیکن
// 	computerScore := 0 //تعداد پیروزی های کامپیوتر
// 	reader := bufio.NewReader(os.Stdin)
// 	clearScreen() //به این وسیله ترمینال را پاک میکنیم
// 	fmt.Println("ROCK,PAPER & SCISSORS")
// 	fmt.Println("---------------------")
// 	fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
// 	fmt.Println("")
// 	for i := 1; i <= 3; i++ { //این حلقه را میزنیم تا بازی سه بار ادامه پیدا کند
// 		fmt.Println()
// 		fmt.Println("Round", i) //تعداد دفعاتی که بازی انجام شده است
// 		fmt.Println("-------")
// 		fmt.Print("Please enter rock,paper or scissors -> ")
// 		playerChoice, _ = reader.ReadString('\n')                    //اون چیزی که کاربر تایپ کرده را میخوانیم
// 		playerChoice = strings.Replace(playerChoice, "\r\n", "", -1) //اون چیزی که کاربر نوشته را مینویسیم
// 		computerValue := rand.Intn(3)                                //اینجا میگیم که اعداد بین یک تا سه باشد
// 		if playerChoice == "rock" {                                  //اینجا شرط میگذاریم که اگر کاربر تایپ کرده بود سنگ یا کاغذ و یا قیچی مساوی با چه مقدار ثابت باشه
// 			playerValue = ROCK
// 		} else if playerChoice == "paper" {
// 			playerValue = PAPER
// 		} else if playerChoice == "scissors" {
// 			playerValue = SCISSORS
// 		} else {
// 			playerValue = -1
// 		}
// 		fmt.Println()
// 		fmt.Println("Player chose", strings.ToUpper(playerChoice)) //اون چیزی که کاربر تایپ کرده رو به صورت حروف بزرگ و به string تبدیل میکنیم

// 		switch computerValue { //این جا هم مقادیری که کامپیوتر به صورت random میگیره رو تبدیل میکنیم به سنگ کاغذ و قیچی
// 		case ROCK:
// 			fmt.Println("Computer Choose ROCK")
// 			break
// 		case PAPER:
// 			fmt.Println("Computer Choose PAPER")
// 			break
// 		case SCISSORS:
// 			fmt.Println("Computer Choose SCISSORS")
// 			break
// 		default:
// 		}
// 		fmt.Println()
// 		if playerValue == computerValue { //اینجا اگر مقدار کاربر با مقدار بازیکن برابر بود مساوی اعلام میکنه و بعد هم یک دور کم میکند
// 			fmt.Println("It is Draw")
// 			i--
// 		} else if playerValue == -1 {//این قسمت کد جایگزین تمام اون کد های دیگر شده است به وسیله علامت % که باقی مانده را می آورد
// 			fmt.Println("Invalid Choice")
// 			i--
// 			}else if playerValue == (computerValue + 1) % 3{
// 				playerScore = PlayerWins(playerScore)
// 			} else {
// 			 computerScore = ComputerWins(computerScore)
// 			}//
// 		}
// 	}
// 	fmt.Println("Final Score")
// 	fmt.Println("-----------")
// 	fmt.Printf("Player %d/3, Computer %d/3", playerScore, computerScore) //اینجا هم تعداد برد بازی کن وکامپیوتر را نشان میدهد
// 	fmt.Println()
// 	if playerScore > computerScore { //اینجا هم بردهای بازی کن و کامپیوتر را مقایسه میکند و برنده را اعلام میکند
// 		fmt.Println("Player Wins Game!")
// 	} else {
// 		fmt.Println("Computer Wins Game!")

//		}
//	}
//
// func clearScreen() { //این همون فعل پاک کردن ترمینال هست
//
//		if strings.Contains(runtime.GOOS, "windows") {
//			// windows
//			cmd := exec.Command("cmd", "/c", "cls")
//			cmd.Stdout = os.Stdout
//			cmd.Run()
//		} else {
//			// linux or mac
//			cmd := exec.Command("clear")
//			cmd.Stdout = os.Stdout
//			cmd.Run()
//		}
//	}
//
// func ComputerWins(score int) int { //این هم تعداد پیروزی های کامپیوتر را در خود نگه میدارد و در پایان هر دور نشان میدهد
//
//		fmt.Println("Computer wins!")
//		return score + 1
//	}
//
// func PlayerWins(score int) int { //این هم تعداد پیروزی های بازی کن را در خود نگه میدارد و در پایان هر دور نشان میدهد
//
//		fmt.Println("Player wins!")
//		return score + 1
//	}
//
// /////////////////////////////////////////////عملیات مک کارتی ///////////////////////////////////////////////////////
// func main() {
// 	a := 12
// 	b := 6
// if b != 0 {
// 	c := dividetwonumber(a, b)
// 	if c == 2 {
// 		fmt.Println("we found a two")
// 	}
// }
// if b!= 0 && dividetwonumber(a, b) != 0{//روش مک کارتی درست کار میکند
// 	fmt.Println("found 2")
// }
// if dividetwonumber(a, b) != 0 && b!= 0{//روش مک کارتی درست کار نمیکند
// 	fmt.Println("found 2")
// }
// }

//	func dividetwonumber(x, y int) int {
//		return x / y
//	}
//
// func main() {//به روش مک کارتی عمل میکند
//
//		a := 12
//		b := 6
//		c , err := dividetwonumber(a, b)//این جا هم داره دو تا متغیر رو یکی به عنوان خروجی عددی و یکی را هم به عنوان خروجی خطا در نظر میگیرد
//		if err != nil {//اگر خطا برابر با خالی نبود بیا و خطا رو چاپ کن
//			fmt.Println(err)
//		}else {
//			if c == 2 {
//				fmt.Println("we found a two")
//			}
//		}
//	}
//
//	 func dividetwonumber(x, y int) (int, error) {//این فعل دارد دو ورودی عددی میگیرد و دو خروجی که یکی عدد و یکی خطا هست میدهد
//		if y == 0 {// ورودی دوم اگر عدد 0 بود بیا و صفر را برگردان و یک خطا چاپ کن
//			return 0, errors.New("cannot divide by 0")
//		}
//	 	return x / y, nil//تقسیم کن دو عدد ورودی را و برای خطا هم خالی برگردان
//	 }
//
// /////////////////////////////////////////////عملگر های انتصاب////////////////////////////////////////////////////
//
//	func main (){
//		x := 10
//		x *= 2 //عملگر انتساب
//		x /= 3 //عملگر انتساب
//		fmt.Println(x)
//	}
//
// //////////////////////////////////////////string یا رشته //////////////////////////////////////////////////////////////////
//
//	func main() {
//		name := "Hello, world!"
//		fmt.Println("string: ", name)
//		for i := 0; i < len(name); i++ { //این مقدار بایت رشته را نشان میدهد
//			fmt.Println("%x ", name[i])
//		}
//		name = "سلام دنیا"
//		fmt.Println("Index\tRune\tString") //این علامت \t فکر کنم یک ستون درست میکند و اطلاعات را درون خود میگذارد
//		for x, y := range name {
//			fmt.Println(x, "\t", y, "\t", string(y)) //اگر خواستیم رشته ای را در هر نوع زبانی بفهمیم چی هست به این روش عمل میکنیم
//		}
//		// سه روش برای این که رشته ها را در کنار هم قرار دهیم
//		h := "Hollo"
//		w := "World"
//		// روش اول
//		mystring := h + w
//		fmt.Println(mystring)
//		//روش دوم
//		mystring = fmt.Sprintf("%s%s", h, w)
//		fmt.Println(mystring)
//		//روش سوم
//		var sb strings.Builder
//		sb.WriteString(h)
//		sb.WriteString(w)
//		fmt.Println(sb.String())
//		// روش دیگر برای نشان دادن تعدادی از رشته ها
//		name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//		fmt.Println(name[0:13]) //این میاد از 0 تا حرف 13 را میخواند
//	}
//
// //////////////////////////////////////////////////چاپ یک رشته به وسیله index آن ها////////////////////////////////////////////////////////////////
// func main() {
// 	courselearn := "سلام دوست من چطوری امروز حالت خوبه"
// 	fmt.Println(string(courselearn[10]))
// 	for i := 0; i < 17; i++ {
// 		fmt.Println(string(courselearn[i]))
// 	}
// 	fmt.Println(len(courselearn)) //این طول رشته رو میاد میشماره
// 	var myslice []string
// 	myslice = append(myslice, "one")
// 	myslice = append(myslice, "two")
// 	myslice = append(myslice, "three")
// 	fmt.Println(len(myslice))              // این یکی میاد و تعداد اعضا این slice رو میگه
// 	fmt.Println(myslice[len(myslice)] - 1) //این یکی میاد و آخرین عضو رو نشان میدهد
// 	for i := 0; i < len(myslice); i++ {    //این میاد و به تعداد اعضای داخل slice این حلقه رو اجرا میکند

// 	}
// 	////////////////////////////////////////////////استفاده از کتابخانه strings //////////////////////////////////////////////////
// 	courses := []string{
// 		"Learn Go Begginers Crash Course",
// 		"Learn Java Begginers Crash Course",
// 		"Learn Python Begginers Crash Course",
// 		"Learn C# Begginers Crash Course",
// 	}
// 	for _, x := range courses {
// 		if strings.Contains(x, "Go") { //این contains میاد واون کلمه ای که در قالب رشته هست رو اگر داخل اون جمله یا کلمه باشد تما اون رو برمیگردونه
// 			fmt.
// 				Println("Go is found in", x, "and Index is", strings.Index(x, "Go")) //این index هم میاد و عددی که اون کلمه داره رو میاره
// 			//خروجی بالا Go is found in Learn Go Begginers Crash Course and Index is 6
// 		}
// 		//                      1         2          3        4
// 		//            0123456789012345678901234567890123456789012345
// 		newString := "Go is a greate programming language.Go for it!"
// 		fmt.Println(strings.HasPrefix(newString, "Go"))     // خروجی این true هست که به ما میگوید این کلمه داخل این هست
// 		fmt.Println(strings.HasPrefix(newString, "Puthon")) // خروجی این false هست که به ما میگوید این کلمه داخل این نیست
// 		fmt.Println(strings.HasSuffix(newString, "!"))      // خروجی این true هست که برای اشکال به کار میرود
// 		fmt.Println(strings.Count(newString, "GO"))         // خروجی این 2 است چون دوبار در جمله ما به کار رفته است و اگر نباشد 0 برمی گرداند
// 		fmt.Println(strings.Index(newString, "GO"))         // خروجی این 0 است چون در اولین خانه قرار دارد ولی اگر در رشته نباشد به ما 1-برمی گرداند
// 		fmt.Println(strings.LastIndex(newString, "GO"))     // خروجی این 36 است چون میرود و از آخر رشته شروع میکند
// 		if strings.Contains(newString, "Go") {
// 			newString = strings.Replace(newString, "GO", "Golang", 1)  //این میاد و کلمه Golang روجایگزین اولین کلمه  Go میکند
// 			newString = strings.Replace(newString, "GO", "Golang", -1) //این خط میاد و کلمه golang رو جایگزین اولین وآخرین کلمه Go به کار رفته در جمله میکند
// 			newString = strings.ReplaceAll(newString, "GO", "Golang")  //این خط میاد و Golang رو جایگزین تمام کلمات Go میکند
// 		}
// 		searchString := strings.ToLower(newString) //به وسیله این تمام جمله بالارا به حروف کوچک تبدیل کردیم
// 		//if strings.Contains(searchString,"go"){ //اگر از کد بالا استفاده کنیم جواب این شرط Find it می شود در غیر این صورت  خیر
// 		//البته به جای این خط بالا میتوانیم به این شکل هم عمل کنیم که کد ما خلاصه تر میشود
// 		if strings.contains(strings.ToLower(newString)) {
// 			fmt.Println(("Find it!"))
// 		} else {
// 			fmt.Println(("Did not Find it!"))
// 		}
// 		fmt.Println(strings.ToLower(newString))                //این تمام کلمات رو بزرگ میکند
// 		fmt.Println(strings.ToUpper(newString))                //این تمام کلمات رو کوچک میکند
// 		fmt.Println(strings.Title(newString))                  // این حرف اول هر کلمه را فقط بزرگ میکند
// 		fmt.Println(strings.Title(strings.ToLower(newString))) //البته به این شکل هم استفاده میشود که خیلی کاربردی است که ابتدا کوچک میشود و سپس حرف اول هر کلمه بزرگ میشود
// 		//در زبان Go شما میتوانید که 2رشته (string)را با هم مقایسه کنیدزیرا آنها را براساس اون عددی که دارند مقایسه میکند
// 		if "Alpha" > "Absolute" {
// 			fmt.Println("Alpha is greater than absolute") //خروجی این گزینه میشود
// 		} else {
// 			fmt.Println("Alpha is not greater than absolute")
// 		}
// 		//برای این که فاصله یا همان space اول و آخر یک کلمه یا یک ایمیل را از بین ببریم از TrimSpace استفاده میکنیم
// 		badEmail := " hosain@gmail.com "
// 		badEmail = strings.TrimSpace(badEmail)
// 		fmt.Printf("=%s=", badEmail)
// 		//خروجی میشود این"=hosain@gmail.com="
// 		//چند مثال در باره ی استفاده از Package Strings در پروژه ها
// 		matches := []string{
// 			"are you",
// 			"what",
// 			"how",
// 			"because",
// 			"sorry",
// 			"i think",
// 			"friend",
// 			"yes",
// 		}
// 		for i := 0; i < len(matches); i++ {
// 			match := matches[i]
// 			position := strings.Index(strings.ToLower(userInput), match)

// 		}

// 	}
// }

// func main() {//در اینجا ما به وسیله فعل replaceNTh توانستیم اون رشته ای که کلمات داخلش شبیه به هم بود را اون یکی که مد نظر داشتیم رو تغییر بدهیم
// 	str := "alpha alpha alpha alpha alpha"
// 	str = replaceNTh(s, "alpha", "Beta", 3)
// 	fmt.Println(str)
// }
// func replaceNTh(s, old, new string, n int) string {
// 	i := 0
// 	for j := 1; j <= n; j++ {
// 		x := strings.Index(s[i:], old)
// 		if x < 0 {
// 			break
// 		}
// 		i = i + x
// 		if j == n {
// 			return s[:i] + new + s[i+len(old):]
// 		}
// 		i+= len(old)
// 	}
// 	return s
// }
///////////////////////////////////////////////نحوه اتصال فایل Golang به مرورگر///////////////////////////////////////////////////////////////
// package main
// func homePage(w http.ResponseWriter, r *http.Request){
// 	html := `<strong>Hello world</strong>`
// 	w.Header().Set("Content-Type", "text/html")
//   fmt.Fprint(w, html)
// }
//  func main (){
//   http.HandleFunc("/", homePage)

//   log.Println("started web server on port 8080")
//   http.ListenAndServe(":8080", nil)
//  }
//بعد از وارد کردن کد های بالا در ترمینال برنامه را اجرا کنید
// go run main.go
//سپس به مرورگر خود بروید و این را سرچ کنید
//http://localhost:8080/
//حالا شما میتوانید از مرورگر استفاده کنید
//البته یک پروژه سنگ کاغذ قیچی هم که در صفحه وب بار گذاری شده هم ساختیم که در یک فایل دیگر هست تا بتوانیم اجرا کنیم
/////////////////////////////////////////////////////فصل دوم UDEMY شروع شد//////////////////////////////////////////////////////
//ما در زبان گولنگ بعضی از function ها هستند که دو خروجی دارند
// func main(){
// 	oneex, twoex := hello()

// 	fmt.Println(one, two)
// }
// func hello () (string, string){
// 	return "one exit", "two exite"
// }
////////////////////////////////////////pointers قسمت 9///////////////////////////////////////
//این را هم در نظر بگیرید که این متغیر داخل یک function تعریف شده است و از قانون scope پیروی میکند و ما داریم آن را دور میزنیم
// func main(){
// 	mystring := "Green"
// 	log.Println("PRINT 1", mystring)
//     pointeringo (&mystring)//آدرس را به این شکل دریافت میکند شما مرجع یک متغیر را بدست می آورید
// 	log.Println("after", mystring)
// }
// func pointeringo(s *string){//ایجا میگه به آدرسی بروید که اینجا دریافت کرده اید ومحتوای آن هر چی هست تغییر بدهید
// 	log.Println("address in ram", *s)//هنگامی که نیاز دارید به یک اشاره گر واقعی مراجعه کنید از ستاره*استفاده کنید
// 	newValue:= "red"
// 	*s = newvalue//این مقدار را داخل این آدرس قرار بده
// }
//////////////////////////////////////////////////struct , نام گذاری متغیر هاقسمت 10////////////////////////////////////////////////
// var s = "seven"//باید در نام گذاری ها دقت کرد زیرا به مشکل میخوریم
// func main(){
// 	var s2 = "six"
// 	log.Println(s)
// 	log.Println(s2)
// 	saysome("xxx")
// }
// func saysome(s3 string) (string, string){//اگر اینجا از s استفاده کنیم به ما xxx را بر میگرداند
// 	log.Println(s)//اینجا برای ما seven رو بر میگردونه
// 	return s3,"word"
// }
// type User struct {//نحوه ی تعریف یک ساختار
// 	FirstName string
// 	LastIndex string
// 	PhoneNumber string
// 	Age int
//     BirthDate time.Time //این برای ذخیره تاریخ هست
// }
////////////////////////////////////////////////اختصاص دادن یک struct به یک function توسط pointer * قسمت 11////////////////////////////////
// //ما یک func تعریف میکنیم و میگوییم که از نوع struct باشد به وسیله pointer به مثال زیر دقت کنید code 1
// func (s *User) saystruct() string {//روش دیگر برای استفاده از struct,pointer
// 	return s.FirstName //code1
// }
// func main(){
// 	user := User{//نحوه ی مقدار دهی به یک ساختار
// 		FirstName: "Hosain",
// 		LastName: "Rezaee",
// 		PhoneNumber: "09387860379",
// 		Age: 30,
// 		BirthDate: 14/9/1379,
// 	}
// 	log.Println(user.FirstName, user.LastName, user.PhoneNumber, user.Age, user.BirthDate)//نحوه ی فراخوانی یک ساختار
// 	log.Println(user.saystruct()) //code 1
// }
////////////////////////////////////////ساخت map قسمت 12///////////////////////////////////////////
// func main (){
// //	              key     value
// 	mymap := make(map[string] string)
// 	mymap2 := make(map[string] int)
// 	mymap3 := make(map[int] string)

// 	//نحوه اضافه کردن به map
// 	mymap["dog"] = "rengo"
// 	mymap["cat"] = "malos"
// 	mymap2["First"] = 3
// 	mymap2["Second"] = 2
// 	// نحوه تغییر مقدار یک map به وسیله key
// 	log.Println(mymap["cat"])
// 	mymap["cat"] = "shangol"
// 	//نحوه فراخوانی یک mapبه وسیله key
// 	log.Println(mymap["dog"])
// 	log.Println(mymap["cat"])
// 	log.Println(mymap2["First"], mymap2["Second"])
// }
//همچنین میتوانیم از map به وسیله struct نیز استفاده کنیم
// type User struct {
// 	FirstName string
// 	LastName string
// }
// func main(){       // int
// 	mymap := make(map[string]User)
// 	me := User{
// 		FirstName:"hosain",
// 		LastName:"rezaee",
// 	} //   "12"
// 	mymap["name"] = me
// 	log.Println(mymap["name"].FirstName)
// }
//نکته :شما هرگز نمی توانید مقادیری که داخل map ذخیره میکنید را مرتب کنید و به همان ترتیبی که وارد کردید
//خارج کنید شما همیشه باید از key ها استفاده کنید
//نحوه ی ساخت یک slice
// func main(){

// 	var myslice []string
// 	var myslice2 []int
// 	//نحوه ی اضافه کردن یک مقدار به slice
// 	myslice = append(myslice,"hosain")
// myslice2 = append(myslice,3)
// myslice2 = append(myslice,1)
// myslice2 = append(myslice,2)
// //برای این که یک slice را مرتب کنیم به ترتیب عدد یا حروغ الفبا از روش زیر استفاده میکنیم
// log.Println(myslice2)
// //خروجی بالا به این شکل است[3,1,2]
// sort.Intn(myslice2)
// //خروجی بالا به این شکل است بعد از مرتب سازی[1,2,3]
// log.Println(myslice)
// //روش خلاصه نویسی هم به این شکل است و ما بیشتر استفاده میکنیم
// numbers := []int{1,2,3,4,5,6,7,8,9}
// names := []string{"cat","fish","dog","car","book"}
// //همان طور میتوانیم فقط قسمتی از یک slice را نشان دهیم
// log.Println(numbers)
// log.Println(numbers[0:3])
// log.Println(numbers[6:9])
// }
////////////////////////////////////////شرط ها قسمت 13//////////////////////////////////////////////////
// func main(){
// mynum := 100
// istrue := false
// if mynum > 99 && !istrue {
// 	log.Println("1")
// 	}else if mynum < 100 && istrue {
// 	log.Println("2")
// 	}else if mynum == 101 || istrue {
// 	log.Println("3")
// 	}else if mynum > 1000 && istrue == false {
// 	log.Println("4")

// }
// myvar := "cat"
// switch myvar {
// case "cat":
// 	log.Println("1")
// case "dog":
// 	log.Println("2")
// case "fish":
// 	log.Println("3")
// default:
// 	log.Println("4")
// }
// }
////////////////////////////////////////////for قسمت 14/////////////////////////////////////////////////////////
// func main(){
// 	animals := []string {"dog","cat","horse","cow"}
// 	for i,animal := range animals {
// 		log.Println(i,animal)
// 		//خروجی حیوانات به همراه عدد هایی کنار آنها زیرا خروجی range همین است یک عدد و یک مقدار برمیگرداند
// 		//0 dog
// 		//1 cat
// 		//2 horse
// 		//3 cow
// 	}
// 	//اگر بخواهیم که از اعداد استفاده نکنیم آن را با یک خط تیره جایگزین میکنیم
// 	for _, animal := range animals {
// 		log.Println(animal)
// 		// dog
// 		// cat
// 		// horse
// 		// cow
// 		}
// 		animals2 := make(map[string]string)
// 		animals2["dog"] = "good"
// 		animals2["cat"] = "bad"
// 		for animaltype, animal := range animals2 {
// 			log.Println(animaltype, animal)
// 		}
// 		//در زبان گولنگ هر رشته (String) در واقع یک بایت (byte)است
// 		var firstLine = "سلام دوستان من حسین هستم"
// 		for i, l := range firstLine {
// 			log.Println(i,":", l)
// 			//خروجی این یک سری عدد هست که مقدار بایت هر کلمه را نشان میدهد وهمچنین رشته ها یک سرس نشانه ها هستند
// 		}
// 		type User struct {//ما در اینجا یک کاربر از نوع struct تعریف کردیم
// 			FirstName string
// 			LastName string
// 			Email string
// 			Age int
// 		}
// 		var users []User//در اینجا هم گفتیم که تمام کاربر هامون از نوع کاربر بالا باشند
// 		users = append(users, User{"john", "Smith","JohnSmith@gmail.com", 30 })
// 		users = append(users, User{"Mary", "jones","Mary@gmail.com", 20 })
// 		users = append(users, User{"sally", "Brown","sallyBrown@gmail.com", 45 })
// 		users = append(users, User{"Alex", "Anderson","AlexAnder@gmail.com", 33 })
// 		for _, l := range users {
// 			log.Println(l.FirstName, l.LastName, l.Email, l.Age)
// 		}
// }
////////////////////////////////////////////Interface قسمت 15/////////////////////////////////////////////////////////
// type Animal interface {
// 	Says() string
// 	NumberOfleg() int
// }
// type Dog struct {
// 	Name  string
// 	Breed string
// }
// type Gorila struct {
// 	Name        string
// 	Color       string
// 	NumberOfleg int
// }

// func main() {
// 	dog := Dog{
// 		Name:  "Samson",
// 		Breed: "German Shephered",
// 	}
// 	PrintInfo(&dog)
// 	gorilla := Gorilla{
// 		Name:          "jack",
// 		Color:         "black",
// 		NumberOfTeeth: 38,
// 	}
// 	PrintInfo(&gorilla)
// }
// func PrintInfo(a Animal) {
// 	fmt.Println("", a.Says(), "", a.NumberOfleg())
// }
// func (d *Dog) Says() string {
// 	return "woof"
// }
// func (d *Dog) NumberOfleg() int {
// 	return 4
// }
// func (g *Gorila) Says() string {
// 	return "Ugh"
// }
// func (g *Gorila) NumberOfleg() int {
// 	return 38
// }
////////////////////////////////////////////Channel قسمت 18/////////////////////////////////////////////////////////
// package helper

// func RandomNumber(n int) int{
// rand.seed(time.Now().UnixNano())
// value := rand.Intn(n)
// return value
// }

// package main

// const numpool = 1000
// func CalculateValue(intchan chan int){
//  randomNumber := helper.RandomNumber(numpool)
//  intchan <- randomNumber
// }
// func main(){
// 	intchan := make(chan int)
// 	defer close(intchan)
// 	go calculateValue(intchan)
// 	num := <-intchan
// 	log.Println(num)
// }
// type Person struct {
// FirstName string `json:"first_name"`
// LastName string `json:"last_name"`
// HairColor string `json:"hair_color"`
// HasDog bool `json:"has_dogs"`
// }
// func main(){
// 	myJson := `
// 	[
// 		{
// 			"first_name":"Hosain",
// 			"last_name":"Rezaee",
// 			"hair_color":"Black",
// 			"has_dogs": true
// 		},
// 		{
// 			"first_name":"Zahra",
// 			"last_name":"Rezaee",
// 			"hair_color":"Black",
// 			"has_dogs": false
// 		}
// 	]`
// 	var unmarshaled []Person//دراین روش ما اطلاعات را به صورت یک slice میبینیم
// 	err:= json.Unmarshal([]byte(myJson), &unmarshaled)
// 	if err != nil {
// 		log.Println("خطا",err)
// 	}
// 	log.Printf("unmarshaled: %v", unmarshaled)//خروجی این قسمت [{zahra Rezaee Black false} ]
// 	// write Json from Struct به این روش ما اطلاعات رو که گرفتیم به شکل myJson میتوانیم بخوانیم خروجی مثل myJson است
// 	var mySlice []Person
// 	var m1 Person
// 	m1.FirstName = "Hamed"
// 	m1.LastName = "Rezaee"
// 	m1.HairColor = "black"
// 	m1.HasDog = false
// 	mySlice = append(mySlice, m1)
// 	var m2 Person
// 	m2.FirstName = "vahid"
// 	m2.LastName = "Rezaee"
// 	m2.HairColor = "white"
// 	m2.HasDog = true
// 	mySlice = append(mySlice, m2)
// 	newJson, err := json.MarshalIndent(mySlice, "", "     ")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(newJson))
// }
/////////////////////////////////////////////////////hello world قسمت 22/////////////////////////////////////////////////////
// func main(){
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		n, err := fmt.Fprintf(w, "hello world")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
// 	})
// 	_ = http.ListenAndServe(":8080", nil)
// }
//////////////////////////////////////////////////مدیریت صفحات وب ومدیریت خطاها قسمت 24و25////////////////////////////////////////////////
package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) { //به این وسیله ما یک صفحه وب داریم که یک عملیات انجام میدهد و یک پاسخ را میفرستد
	fmt.Fprintf(w, "This is a Home page")
}
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is About page and 2+2 is %d", sum))
}
func addValue(x, y int) int { //این یک function است که یک عملیات جمع را انجام میدهد واز آن در صفحات وب استفاده شده است
	return x + y
}
func divide(w http.ResponseWriter, r *http.Request) {
	f, err := DivideValue(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 0.0, f))
}
func DivideValue(x, y float32) (float32, error) { // ما به این وسیله خطاهایی که هست رو کنترل میکنیم وخروجی این فعل یک عدد اعشاری و یک خطا هست
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
func main() {
	http.HandleFunc("/", Home) //در این قسمت هم ما اون functionکه یک صفحه وب رو میسازد صدا میزنیم وبه اون یک پسوند میدهیم تا به وسیله آن صفحه وب را پیدا کنیم
	http.HandleFunc("/About", About)
	http.HandleFunc("/divide", divide)
	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
//////////////////////////////////////////////////ساخت صفحات html////////////////////////////////////////////////
package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) { //به این وسیله ما یک صفحه وب داریم که یک عملیات انجام میدهد و یک پاسخ را میفرستد
	renderTemplate(w, "home.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
func renderTemplate(w http.ResponseWriter, tmpl string){
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template",err)
		return
	}
}


func main() {
	http.HandleFunc("/", Home) //در این قسمت هم ما اون functionکه یک صفحه وب رو میسازد صدا میزنیم وبه اون یک پسوند میدهیم تا به وسیله آن صفحه وب را پیدا کنیم
	http.HandleFunc("/About", About)
	
	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
















































































































































































































































































































