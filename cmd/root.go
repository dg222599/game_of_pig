// /*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>
// */
// package cmd

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/spf13/cobra"
// )

// func HandleError(err error) {
// 	fmt.Println("Exiting...Encountered error in -->" , err.Error())
// 	os.Exit(1)
// }

// var (
// 	flag1,flag2,totalGames,player1Start,player2Start,player1End,player2End int
// 	resultType,resultSize string
// )

// // rootCmd represents the base command when called without any subcommands
// var rootCmd = &cobra.Command{
// 		Use:   "game_of_pig hold1 hold2",
// 		Short: "It is a CLI based game involving 2 players.",
// 		Long: `A longer description that spans multiple lines and likely contains
// 				examples and usage of using your application. For example:

// 				Cobra is a CLI library for Go that empowers applications.
// 				This application is a tool to generate the needed files
// 				to quickly create a Cobra application.`,
// 	// Uncomment the following line if your bare application
// 	// has an action associated with it:
// 	 	Run: func(cmd *cobra.Command, args []string) {

// 		player1strat,err  := strconv.Atoi(args[0])
// 		HandleError(err)

// 		player2strat,err  := strconv.Atoi(args[1])
// 		HandleError(err)

// 		player1flag,err := cmd.Flags().GetInt("f1")
// 		HandleError(err)

// 		player2flag,err := cmd.Flags().GetInt("f2")
// 		HandleError(err)

// 		resultsflag,err := cmd.Flags().GetString("v")
// 		HandleError(err)

// 		resultSize,err := cmd.Flags().GetString("s")
// 		HandleError(err)

// 		// var player1score,player2score,currentTurnScore,currentStrat,player1Wins,player2Wins,currentPlayer int64
// 		var winPct,lossPct float64
// 		var limit1,limit2 int
// 		var shortResultFile,longResultFile *os.File

// 		limit1 = 1
// 		limit2 = 1

// 		if player1flag !=1{
// 			 if player1flag > player1strat {
// 				 limit1 = player1flag
// 			 } else {

// 				fmt.Printf("Player 1 will hold only till %d no point in giving %d since it is lesser than %d\n\n",player1strat,player1flag,player1strat)
// 				limit1 = player1strat
// 			 }
// 		}

// 		if player2flag !=1{
// 			if player2flag > player2strat {
// 				limit2 = player2flag
// 			} else {
// 				fmt.Printf("Player 2 will hold only till %d no point in giving %d since it is lesser than %d\n\n",player2strat,player2flag,player2strat)
// 				limit2 = player2strat
// 			}
// 		}

// 		var fileOpenError error

// 		if resultType != "" {
// 			longResultFile,fileOpenError = os.Create("long_result.txt")
// 			if fileOpenError!=nil{
// 				fmt.Println("error in opening the long file-->",fileOpenError.Error())
// 				os.Exit(1)
// 			} else {
// 				fmt.Println("Saving results in long format file long_result.txt")
// 			}
// 			defer longResultFile.Close()
// 		}

// 		if resultSize !=""{
// 			shortResultFile,fileOpenError = os.Create("short_result.txt")
// 			if fileOpenError!=nil{
// 				fmt.Println("error in opening the short file-->",fileOpenError.Error())
// 				os.Exit(1)
// 			} else {
// 				fmt.Println("Saving results in short format file short_result.txt")
// 			}
// 			defer shortResultFile.Close()
// 		}

// 		totalGames = (limit1 -1)*limit2

// 		for player1times:=1;player1times<=limit1;player1times++ {
// 				player1Wins = 0
// 				player2Wins = 0
// 			for player2times:=1;player2times<=limit2;player2times++{

// 				player1strat = player1times
// 				player2strat = player2times
// 				if player1strat == player2strat {
// 					continue
// 				}
// 				for gameNumber:=0;gameNumber<10;gameNumber++{

// 					player1score = 0
// 					player2score = 0

// 					currentTurnScore = 0
// 					currentStrat = int64(player1strat)
// 					currentPlayer = 1

// 					//will run as long as any one of the  player wins
// 					for ; ; {

// 						 //rolling the dice
// 						 rand.NewSource(time.Now().UnixNano())
// 						 diceNumber := rand.Intn(6) + 1

// 						 if diceNumber == 1 {
// 						   //player change happens now
// 						   if currentPlayer == 1 {
// 							   currentPlayer = 2
// 							   currentStrat = int64(player2strat)
// 						   } else {
// 							   currentPlayer = 1
// 							   currentStrat = int64(player1strat)
// 						   }

// 						   currentTurnScore = 0

// 						 } else {

// 							 currentTurnScore+=int64(diceNumber)

// 							 if currentTurnScore>=currentStrat {

// 								  if currentPlayer == 1 {
// 									  player1score+=currentTurnScore
// 									  currentPlayer = 2
// 									  currentStrat = int64(player2strat)
// 									  currentTurnScore = 0
// 									  if player1score >=100 {
// 										   break
// 									  }
// 								  } else {
// 									  player2score+=currentTurnScore
// 									  currentPlayer = 1
// 									  currentStrat = int64(player1strat)
// 									  currentTurnScore = 0
// 									  if player2score >=100 {
// 										   break
// 									  }
// 								  }

// 							 }

// 						 }

// 					}

// 					if player1score >= 100 {
// 					   player1Wins++
// 					} else {
// 					   player2Wins++
// 					}
// 			  	}

// 				if resultsflag != "" && resultSize!="" {
// 					winPct = (float64(player1Wins)/10)*100
// 					lossPct = (float64(player2Wins)/10)*100

// 					_,err :=fmt.Fprintf(longResultFile,"Holding at %d vs Holding at %d: wins: %d/10 (%f%%), losses: %d/10 (%f%%)\n\n",player1strat,player2strat,
// 					player1Wins,winPct,player2Wins,lossPct)
// 					if err!=nil{
// 						fmt.Println("error in writing the results in the file")
// 						os.Exit(1)
// 					}
// 				}

// 			}

// 			if resultSize != "" {

// 				winPct = (float64(player1Wins)/float64(totalGames))*100
// 				lossPct = (float64(player2Wins)/float64(totalGames))*100

// 				_,err :=fmt.Fprintf(shortResultFile,"Result: Wins, losses staying at k = %d: %d/%d (%f%%), %d/%d (%f%%)\n\n",player1strat,player1Wins,totalGames,winPct,player2strat,
// 						totalGames,lossPct)
// 				if err!=nil{
// 					fmt.Println("error in writing the results in the file")
// 					os.Exit(1)
// 				}

// 			}

// 		}
// 	},
// }

// // Execute adds all child commands to the root command and sets flags appropriately.
// // This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() {
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }

// func init() {
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.

// 	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.game_of_pig.yaml)")

// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// 	rootCmd.Flags().IntVar(&flag1,"f1",1,"Used for player 1 game type")
// 	rootCmd.Flags().IntVar(&flag2,"f2",1,"Used for player 2 game type")
// 	rootCmd.Flags().StringVar(&resultType,"v","","Used for specifying the results display format")
// 	rootCmd.Flags().StringVar(&resultSize,"s","","Used for choosing the results size for display")
// }

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
	Use:   "game_of_pig hold1 hold2",
	Short: "It is a CLI based game involving 2 players.",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntVar(&flag1,"f1",1,"Used for player 1 game type")
	rootCmd.Flags().IntVar(&flag2,"f2",1,"Used for player 2 game type")
	rootCmd.Flags().StringVar(&longResult,"v","","Used for specifying result format in long")
	rootCmd.Flags().StringVar(&shortResult,"s","","Used for specifying the result format in short")
}




