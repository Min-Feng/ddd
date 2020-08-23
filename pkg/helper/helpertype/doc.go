// package helpertype 主要有兩個功能
//
// 1. 提供額外的函數給 基礎型別 使用,
// 函數名稱開頭, 會標明此函數所支援的型別
//
// 2. 定義新的型別, 取代基礎型別, 比如 Time
//
// 此套件對特定型別通常只有少量的支援函數
// 若函數過多, 需要思考重構, 如何切分功能到新的 package
package helpertype