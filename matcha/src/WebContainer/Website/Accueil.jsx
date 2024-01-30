import React from 'react';
import Box from '@mui/material/Box';
import BackImg from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MainPage/backgroung-unsplash.jpg';

const Accueil = () => {
	return (
		<Box className="App__WebContainer__Website__Main">
			<img src={BackImg} style={{ width:"100%", height:"80%"}}/>
		</Box>
	);
}

export default Accueil;