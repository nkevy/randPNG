package main
import(
	"fmt"
	"image"
	"image/color"
//	"image/draw"
	"image/png"
	"math/rand"
	"strings"
	"time"
	"os"
 //      "errors"
)
func validname(name string) bool{
	comp := []string{"?","%","*",":","|","\\","/"}
	for _,c := range comp{
		if strings.Contains(name,c){
			return false;
		}
	}
	return true;
}
func randColor() color.RGBA{
	seed := time.Now().UnixNano()
        rand.Seed(seed)
	fmt.Println(seed)
        //clr := rand.Intn(255)
	//off := float64(clr)/float64(255)
	r :=uint8(rand.Intn(255))
	g :=uint8(rand.Intn(255))
	b :=uint8(rand.Intn(255))
	a :=uint8(rand.Intn(255))
	fmt.Printf("%d %d %d %d\n",r,g,b,a)
        return color.RGBA{r,g,b,a}
}
func ofc(c int) uint8{
	if rand.Intn(6)>3{
		return uint8(c-rand.Intn(255)%10)
	}else{
		return uint8(c+rand.Intn(255)%10)
	}
}
func offColor() color.RGBA{
	seed := time.Now().UnixNano()
        rand.Seed(seed)
	fmt.Println(seed)
	cent := 200
        //clr := rand.Intn(255)
	//off := float64(clr)/float64(255)
	r :=ofc(cent)
	g :=ofc(cent)
	b :=ofc(cent)
	a :=ofc(cent)
	fmt.Printf("%d %d %d %d\n",r,g,b,a)
        return color.RGBA{r,g,b,a}
}
func fill(w,h int) image.Image{
	color := randColor()
	org := image.Point{X:0,Y:0}
	max := image.Point{X:w,Y:h}
	img := image.NewRGBA(image.Rectangle{org,max})
	for i:=0;i<w;i++{
		for j:=0;j<h;j++{
			img.Set(i,j,color)
		}
	}
	return img
}
func hatch(w,h int) image.Image{
	c1 := offColor()
	c2 := offColor()
	c3 := offColor()
	org := image.Point{0,0}
	max := image.Point{w,h}
	img := image.NewRGBA(image.Rectangle{org,max})
	for i:=0;i<w;i++{
		for j:=0;j<h;j++{
			if i%2==0 && j%2==0{
				img.Set(i,j,c1)
			}else if i%3==0 && j%3==0{
				img.Set(i,j,c2)
			}else{
				img.Set(i,j,c3)
			}
		}
	}
	return img
}
func cross(w,h int) image.Image{
	c1 := randColor()
	c2 := randColor()
	org := image.Point{0,0}
	max := image.Point{w,h}
	img := image.NewRGBA(image.Rectangle{org,max})
	for i:=0;i<w;i++{
		for j:=0;j<h;j++{
			if i%2==0 && j%2==0{
				img.Set(i,j,c1)
			}else{
				img.Set(i,j,c2)
			}
		}
	}
	return img
}
func makePNG(fname string) bool{
	filename := fname + ".png"
	f,err := os.OpenFile(filename,os.O_WRONLY|os.O_CREATE,755)
	defer f.Close()
	if err!=nil{
		return false
	}
	// 16:9 for phones is 1920x1080 or 1280x720
	width:=720
	height:=1280
	png.Encode(f,hatch(width,height))
	return true
}
func main(){
	if len(os.Args)!=2{
		fmt.Println("usage: randPNG <filename>")
		return
	}
	fname := os.Args[1:][0]
	if !validname(fname){
		fmt.Println("usage: randPNG <filename> , your file name was invalid")
		return
	}
	if !makePNG(fname){
		fmt.Println("creation of png failed check the runtime permitions")
	}
}
