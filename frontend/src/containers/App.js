import React from 'react';
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import { createTheme, ThemeProvider } from '@mui/material/styles';

import Index from './routes'
import Login from "./routes/login";

const defaultTheme = createTheme();

function App() {
    return (
      <ThemeProvider theme={defaultTheme}>
          <div className="App">
              <Router>
                  <Routes>
                      <Route path="/*" element={<Index />}/>
                      <Route path="/login" element={<Login />}/>
                  </Routes>
              </Router >
          </div>
      </ThemeProvider>
  );
}

export default App