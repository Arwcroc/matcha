import React from 'react';
import Box from '@mui/material/Box';
import BackImg from '../../Images/MainPage/backgroung-unsplash.jpg';

const Accueil = () => {
	return (
		<div>
			<Box className="App__WebContainer__Website__Main">
				<img className="App__WebContainer__Website__Main__Background" src={BackImg}/>
				<Box className="App__WebContainer__Website__Main__PrimarySentence">
					He gonna rolling mechanics
				</Box>
				<Box className="App__WebContainer__Website__Main__CreateAccount">
					Create Account
				</Box>
			</Box>
		</div>
	);
}

export default Accueil;