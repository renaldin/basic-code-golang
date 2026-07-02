package main

type Transaction struct {
	Date   string `json:"date"`
	Amount int    `json:"amount"`
}

func groupTrxDate(mp []Transaction) map[string]int {
	result := make(map[string]int)
	for _, val := range mp {
		result[val.Date] += val.Amount
	}
	return result
}

// func main() {
// 	mp := []Transaction{
// 		{
// 			Date:   "2026-06-20",
// 			Amount: 100,
// 		},
// 		{
// 			Date:   "2026-06-20",
// 			Amount: 50,
// 		},
// 		{
// 			Date:   "2026-06-21",
// 			Amount: 75,
// 		},
// 	}
// 	result := groupTrxDate(mp)
// 	fmt.Println(result)
// }

// 4. Group Transactions by Date

// Problem:
// Given a list of transactions, group them by date and calculate the total amount for each date.

// Input:

// [
//   {
//     "date": "2026-06-20",
//     "amount": 100
//   },
//   {
//     "date": "2026-06-20",
//     "amount": 50
//   },
//   {
//     "date": "2026-06-21",
//     "amount": 75
//   }
// ]

// Expected Output:

// {
//   "2026-06-20": 150,
//   "2026-06-21": 75
// }