import java.math.BigDecimal;

class Calculation extends Thread{
    private int num;
    private int n;
    private BigDecimal result;

    public Calculation(int num, int n, BigDecimal result){
        this.num = num;
        this.n = n;
        this.result = result;
    }
    public void run()
    {
        try {
            // Displaying the thread that is running
            result = new BigDecimal((1+Math.sqrt(5)*num)/2).pow(n);
            System.out.println("done");
        }
        catch (Exception e) {
            // Throwing an exception
            System.out.println("Exception is caught");
        }
    }
}

 
// Main Class
public class FibThr{
    public static void main(String[] args)
    {

        BigDecimal firstBig = new BigDecimal("0");
        BigDecimal secondBig = new BigDecimal("0");
        var first = new Calculation(1,100, firstBig);
        var second = new Calculation(-1, 100, secondBig);
        first.start();
        second.start();
        while(true){
            if(firstBig.compareTo(new BigDecimal("0")) != 0 && secondBig.compareTo(new BigDecimal("0")) != 0){
                System.out.println(new BigDecimal(1/Math.sqrt(5)).multiply(firstBig.subtract(secondBig)).toBigInteger());
                return;
            }
        }
        }
    }

