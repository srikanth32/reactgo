import React from 'react';
import {render} from 'react-dom';


export class App extends React.Component{
render(){
return (
 <div>
 <form method="post" action="/success">
 Username:  <input type="text" name="name"/>

 <button type="button" action="submit">Click Me!</button>
</form>
 </div>
);
}
}
render(<App/>,document.getElementById("app"));
