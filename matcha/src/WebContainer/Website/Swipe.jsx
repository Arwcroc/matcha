import React from 'react';
import Box from '@mui/material/Box';
import BackImg from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MainPage/backgroung-unsplash.jpg';
import Card from '@mui/material/Card';

const Swipe = () => {
	return (
		<div>
			<Box className="App__WebContainer__Website__Main">
				<img className="App__WebContainer__Website__Main__Background" src={BackImg} style={{ width:"100%"}}/>
				<Card className="App__WebContainer__Website__Main__PrimaryCard">
					Je suis une carte
				</Card>
				<Card className="App__WebContainer__Website__Main__OldCard">
					Je suis une vieille carte
				</Card>
				<Card className="App__WebContainer__Website__Main__OldOldCard">
					Je suis une tres vieille carte
				</Card>
				<Card className="App__WebContainer__Website__Main__NextCard">
					Je suis une la prochaine carte
				</Card>
				<Card className="App__WebContainer__Website__Main__NextNextCard">
					Je suis une la prochaine prochaine carte
				</Card>
			</Box>
		</div>
	);
}

export default Swipe;