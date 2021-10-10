import java.math.BigDecimal;
import java.math.BigInteger;
import java.lang.Math;

public class Fib{
    BigInteger nthFib(int n){
        var first = new BigDecimal((1+Math.sqrt(5))/2).pow(n);
        var second = new BigDecimal((1-Math.sqrt(5))/2).pow(n);
        return new BigDecimal(1/Math.sqrt(5)).multiply(first.subtract(second)).toBigInteger();
        
    }
    public static void main(String[] args) {
        Fib f = new Fib();
        System.out.println(f.nthFib(1000000).toString().length());
    }
}