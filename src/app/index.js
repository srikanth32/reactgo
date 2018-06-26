import React from 'react';
import {render} from 'react-dom';
import $ from 'jquery';
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

var URL = 'http://localhost:3033/';

export class App extends React.Component{
     registerData(){
        
        var obj = new Object();
       
    obj["name"] = $("#name").val();
    
    console.log(obj)
    $.ajax({
        url: URL + '/index',
        type: 'GET',
        data: { json: JSON.stringify(obj) },
        cache: false,
        success: function (response) {
            response = JSON.parse(response);
            console.log(response)
            if (response.Status === 'true') {
                alert(response.Message);
            } else {
                alert(response.Message);
            }
        },
        error: function () {
            alert('Unable to update job details !!!');
        },
        complete: function () {
            //self.container.dataLoader('hide');
        }
    });
    }
render(){
return (
    
   
 <div>
 <form >
 Username:  <input type="text" id="name" name="name"/>

 <button type="button"  onClick={(e) => this.registerData(e)}>Click Me!</button>
</form>
 </div>
);
}
}
render(<App/>,document.getElementById("app"));
