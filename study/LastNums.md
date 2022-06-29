# reslice記憶體問題

在原有的slice進行reslice。因為底層的array不會變化，記憶體會一直佔用，直到沒有變數引用該數組。因此很可能出現這麼一種情況，原slice由大量的元素構成，但是我們在原slice的基礎上reslice，雖然只使用了很小一段，但底層數組在記憶體中仍然佔據了大量空間，得不到釋放。比較推薦的做法，使用`copy`替代`re-slice`。


* `generateWithCap`用於隨機生成n 個int 整數，64位機器上，一個int 佔8 Byte，128 * 1024 個整數恰好佔據1 MB 的空間。
* `printMem`用於打印程序運行時佔用的內存大小。




#### 執行`go test -run=^TestLastChars -v `之結論

結果差異非常明顯，lastNumsBySlice耗費了100.14 MB 記憶體，也就是說，申請的100 個1 MB 大小的 記憶體沒有被回收。因為slice雖然只使用了最後2 個元素，但是因為與原來1M 的切片引用了相同的底層數組，底層Array得不到釋放，因此，最終100 MB 的記憶體始終得不到釋放。而 lastNumsByCopy 僅消耗了3.14 MB 的 記憶體。這是因為，通過copy，指向了一個新的底層數組，當origin 不再被引用後，內存會被垃圾回收(garbage collector, GC)。