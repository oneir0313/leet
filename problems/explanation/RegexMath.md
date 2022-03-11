## DP solution with detailed Explanation
1. If `p.charAt(j) == s.charAt(i)`:

   `dp[i][j] = dp[i-1][j-1];`
2. If `p.charAt(j) == '.'`:
    
    `dp[i][j] = dp[i-1][j-1];`
3. `If p.charAt(j) == '*'`: 
   
   here are two sub conditions:
               
    1. if `p.charAt(j-1) != s.charAt(i)`: 
    
        `dp[i][j] = dp[i][j-2]`  
        > in this case, a* only counts as empty

    2. if `p.charAt(i-1) == s.charAt(i) or p.charAt(i-1) == '.'`:
        
        `dp[i][j] = dp[i-1][j]`    
        > in this case, a* counts as multiple a

        or dp[i][j] = dp[i][j-1]

        > in this case, a* counts as single a
                            
        or dp[i][j] = dp[i][j-2] 
        
        >in this case, a* counts as empty