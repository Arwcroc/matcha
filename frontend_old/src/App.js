import React from 'react';
import WebContainer from './WebContainer/WebContainer.jsx'
import { Route, Routes } from "react-router-dom";

function App() {
  return (
	<div className="App">
		<Routes>
			<Route path="/*" element={<WebContainer/>}/>
		</Routes>
    </div>
  );
}

export default App;
