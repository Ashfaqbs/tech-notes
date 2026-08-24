package com.ashfaq.client;

import com.ashfaq.sdk.ApiResponse;
import com.ashfaq.sdk.SdkClient;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
        System.out.println( "Hello World!" );
        SdkClient sdkClient = new SdkClient();
        ApiResponse response = sdkClient.getApiResponse("https://jsonplaceholder.typicode.com/todos/1");
        
        System.out.println("Received API Response: " + response);
        /*
         * C:\tmp\sdk> cmd /C "C:\openjdk\graalvm-jdk-21.0.4+8.1\bin\java.exe @C:\Users\ashfa\AppData\Local\Temp\cp_631y7p4p4db0ruybeybzy66ok.argfile com.ashfaq.client.App "
Hello World!
Received API Response: ApiResponse{userId=1, id=1, title='delectus aut autem', completed=false}

         */
    }
}
