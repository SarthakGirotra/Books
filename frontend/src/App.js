import React, { useLayoutEffect } from 'react'
import './App.css';
import Signupscreen from "./components/login/signup";
import Homepage from "./components/homepage/homepage"
import { useSelector, useDispatch } from "react-redux"
import { home } from "./actions/user"


function App() {
  /*eslint-disable */
  useLayoutEffect(() => {
    dispatch(home())
  }, [])
  /*eslint-enable */

  const dispatch = useDispatch();
  const id = useSelector((state) => state.Auth.id)
  return (
    <div className="App">
      {id ? (<Homepage />) : (<Signupscreen />)}

    </div>
  );
}

export default App;
