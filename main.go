package main

import "fmt"

func main(){

  i := 20

  fmt.Printf("O número é: %v\n", i)

  fat := fatorial(i)

  fmt.Printf("Seu fatorial é: %v\n", fat)
  fmt.Printf(paridade(fat), fat)
  fmt.Printf("O %vº termo da sequencia fibonacci é: %v\n", i, fibonacci(i))

}

func paridade(i int)(p string){
  if i % 2 == 0 {
    return "O número %v é par\n"
  } else {
    return "O número %v é impar\n"
  }
}

func fatorial(i int)(fat int){
  if i == 2 {
    return i
  }

  return i * fatorial(i-1)
}

func fibonacci(n int)(fibo int){
  atual, anterior := 0, 0

  for i := 1; i < n; i++ {
    if(i == 1){
      atual = 1
      anterior = 0
    } else{
      atual = atual +  anterior
      anterior = atual - anterior
    }
  }
  return atual
}
