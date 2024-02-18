/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	longResult,shortResult string
	player1Start,player1End,player2Start,player2End,numberOfGames,currentSetWins1,currentSetWins2,currStrategy,strategy1,strategy2,currScore,player1score,player2score,player1Wins,player2Wins,flag2,flag1 int
	winPct,lossPct float64
	longResultFile,shortResultFile *os.File
	err error
)

func HandleError(err error,msg string) {
	fmt.Println(msg)
	fmt.Println("Exiting...Encountered error in -->" , err.Error())
	os.Exit(1)
}

var rootCmd = &cobra.Command{
	Use:   "game_of_pig strategy1 strategy2 --f1=x --f2=y --v='y' --s='y'",
	Short: "It is a CLI based game involving 2 players.",
	Long: `The Game involves 2 players playing against each other both being computer in this case
.Each player has a strategy to hold until X , where X is the total score they achieve
by successive roll of dice , therefore eac player holds until the total of these rolls crosses 
their hold strategy or they get a 1 , in case they get a 1 they lose all accumulated sum from
that chance and the chance goes to next player. For each pair of strategies a set of 10 games is played.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
	 Run: func(cmd *cobra.Command, args []string) {
		  
		  player1Start,err = strconv.Atoi(args[0]) 
			if err !=nil{
			HandleError(err,"Argument 1 (Player 1 Strategy not given)")
		  }

		  player2Start,err = strconv.Atoi(args[1])
		  if err !=nil{
			HandleError(err,"Argument 2 (Player 2 Strategy not given)")
		  }

		  player1End,err = cmd.Flags().GetInt("f1")
		  if err!=nil{
			HandleError(err,"Flag 1 (Player 1 Range wrongly given)")
		  }

		  player2End,err = cmd.Flags().GetInt("f2") 
		  if err!=nil{
			HandleError(err,"Flag 2 (Player 2 Range wrongly given)")
		  }

		  longResult,err = cmd.Flags().GetString("v")
		  if err!=nil{
			HandleError(err,"Flag for saving long results not given correctly")
		  }

		  shortResult,err = cmd.Flags().GetString("s") 
		  if err !=nil{
			HandleError(err,"Flag for saving short results not given correctly")
		  }

		  if shortResult == "" {
			longResult = "yes" //default setting for printing results
		  } 
		  
		  if longResult != ""{
			longResultFile,err = os.Create("long_result.txt")
			if err!=nil{
				HandleError(err,"Error in Opening the long result text file")
			}
			defer longResultFile.Close()
		  }

		  if shortResult != "" {
			shortResultFile,err = os.Create("short_result.txt")
			if err!=nil{
				HandleError(err,"Error in opening the short result text file")
			}
			defer shortResultFile.Close()
		  }

		  if player1End == 1 {
			//not given flag
			player1End = player1Start
		  } else {
			if player1End <player1Start{
				player1Start = 1
			}
		  }

		  if player2End == 1 {
			player2End = player2Start
		  } else {
			 if player2End < player2Start {
				player2Start = 1
			 }
		  }

		  fmt.Println("Starting the Game for following Strategies")
		  fmt.Printf("Player 1 Holds from %d to %d\n",player1Start,player1End)
		  fmt.Printf("Player 2 Holds from %d to %d\n",player2Start,player2End)

		  if longResult != "" {
			fmt.Println("Long results will be printed in long_result.txt")
		  }
		  if shortResult != "" {
			fmt.Println("Short results will be printed in short_result.txt")
		  }

		  //numberOfGames = (player1End - player1Start+1)*(player2End - player2Start)

		  for strategy1 = player1Start ; strategy1 <= player1End ; strategy1++ {
			 currentSetWins1 = 0
			 currentSetWins2 = 0
			 numberOfGames=0
			 for strategy2=player2Start;strategy2 <= player2End ; strategy2++ {
				if strategy1 == strategy2 {
					continue
				}

				player1Wins = 0
				player2Wins = 0


				for game:=1;game<=10;game++{
					currStrategy = strategy1
					currScore = 0

					player1score = 0
					player2score = 0
					
					for ; ; {
						//rolling the dice
						rand.NewSource(time.Now().UnixNano())
						diceNumber := rand.Intn(6) + 1

						if diceNumber == 1 {

							currScore = 0
							if currStrategy == strategy1{
								currStrategy = strategy2
							} else {
								currStrategy = strategy1
							}
						} else {
							currScore+=diceNumber

							if currScore >= currStrategy {
								if currStrategy==strategy1{
									player1score +=currScore
									currStrategy = strategy2
									if player1score >=100 {
										player1Wins++
										break
									}
								} else {
									player2score+=currScore
									currStrategy = strategy1
									if player2score >=100 {
										player2Wins++
										break
									}
								}
								currScore = 0
							}
						}
					}
				}

				numberOfGames+=10
				currentSetWins1+=player1Wins
				currentSetWins2+=player2Wins
				

				if longResult != ""{
					winPct = (float64(player1Wins)/10)*100
					lossPct = (float64(player2Wins)/10)*100
					
					_,err :=fmt.Fprintf(longResultFile,"Holding at %d vs Holding at %d: wins: %d/10 (%f%%), losses: %d/10 (%f%%)\n\n",strategy1,strategy2,
					player1Wins,winPct,player2Wins,lossPct)
					if err!=nil{
						fmt.Println("error in writing the results in the file")
						os.Exit(1)
					}
				}
			}

			if shortResult != ""{
				winPct = (float64(currentSetWins1)/float64(numberOfGames))*100
				lossPct = (float64(currentSetWins2)/float64(numberOfGames))*100
				
				_,err :=fmt.Fprintf(shortResultFile,"Result: Wins, losses staying at k = %d: %d/%d (%f%%), %d/%d (%f%%)\n\n",strategy1,currentSetWins1,numberOfGames,winPct,currentSetWins2,
						numberOfGames,lossPct)
				if err!=nil{
					fmt.Println("error in writing the results in the file")
					os.Exit(1)
				}
			}


		  }
		 

		  
	 } ,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.game_of_pig.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	
	rootCmd.Flags().IntVar(&flag1,"f1",1,"Used to specify the end limit if player1 needs to hold on different strategies.")
	rootCmd.Flags().IntVar(&flag2,"f2",1,"Used to specify the end limit if player2 needs to hold on different strategies.")
	rootCmd.Flags().StringVar(&longResult,"v","yes","Used to specify if the result needed in set wise format -long results(set to yes by default)")
	rootCmd.Flags().StringVar(&shortResult,"s","","Used to specifying if the result needed in single strategy wise format -  short results , need non empty string to get short foramt results")
}




