package main
import(
	"fmt"
	"image"
	"image/color"
//	"image/draw"
	"github.com/llgcode/draw2d/draw2dimg"
//	"image/png"
	"math/rand"
	"strings"
	"time"
	"os"
 //      "errors"
)
type triangle struct{
	//angle
	points [3]image.Point
}
func  new_triangle() *triangle{
	var tri triangle
	tri.points[0]=image.Point{0,0}
	tri.points[1]=image.Point{0,0}
	tri.points[2]=image.Point{0,0}
	return &tri
}
func (tri triangle) set(x,y int){
		tri.points[0] = image.Point{x,y}
		yoff := rand.Intn(1280)
		xoff := rand.Intn(720)
		p2x := (x+xoff)%720
		p2y := (y+yoff)%1280
		p3x := (p2x-xoff)%720
		p3y := (p2y-yoff)%1280
		tri.points[1] = image.Point{p2x,p2y}
		tri.points[2] = image.Point{p3x,p3y}
}
func getImage(width,height int) *image.RGBA{
	org := image.Point{X:0,Y:0}
	max := image.Point{X:width,Y:height}
	img := image.NewRGBA(image.Rectangle{org,max})
	return img
}
func triFill(w,h int, fname string) bool{
	//color also seeds rand
	color := randColor()
	bcolor := randColor()
	fcolor := randColor()
	img := getImage(w,h)
	//draw2d for draw triangle side length
	gcon := draw2dimg.NewGraphicContext(img)
	gcon.SetStrokeColor(bcolor)
	gcon.SetFillColor(fcolor)
	gcon.SetLineWidth(float64(rand.Intn(3)+3))
	//triangle def
	tri := new_triangle()
	for i:=w;i>-1;i--{
		//fill
		for j:=0;j<h;j++{
			// 2 and 113 for tri
			if (j%113)==0 && (i%113)==0 {
				img.Set(i,j,bcolor)
			}else if 0==(i%3) {
				img.Set(i,j,color)
			}else if 0==(j%9){
				img.Set(i,j,color)
			}else{
				// set triangle points (x,y,width,height) 
				tri.set(i,j)
				gcon.MoveTo(float64(i),float64(j))
				gcon.LineTo(float64(tri.points[1].X),float64(tri.points[1].Y))
				gcon.MoveTo(float64(tri.points[1].X),float64(tri.points[1].Y))
				gcon.LineTo(float64(tri.points[2].Y),float64(tri.points[2].Y))
				gcon.MoveTo(float64(tri.points[2].X),float64(tri.points[2].Y))
				gcon.LineTo(float64(i),float64(j))
				gcon.FillStroke()
				gcon.Close()

			}
			//boarder
			if 0==i || 0==j || w==i || h==j{
				img.Set(i,j,bcolor)
			}
		}
	}
	draw2dimg.SaveToPngFile(fname,img)
	return true
}
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
        //clr := rand.Intn(255)
	//off := float64(clr)/float64(255)
	r :=uint8(rand.Intn(255))
	g :=uint8(rand.Intn(255))
	b :=uint8(rand.Intn(255))
	a :=uint8(rand.Intn(255))
	//fmt.Printf("%d %d %d %d\n",r,g,b,a)
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
	//fmt.Printf("%d %d %d %d\n",r,g,b,a)
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
	//f,err := os.OpenFile(filename,os.O_WRONLY|os.O_CREATE,755)
	//defer f.Close()
	//if err!=nil{
	//	return false
	//}
	// 16:9 for phones is 1920x1080 or 1280x720
	width:=720
	height:=1280
	return triFill(width,height,filename)
	//png.Encode(f,hatch(width,height))
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
