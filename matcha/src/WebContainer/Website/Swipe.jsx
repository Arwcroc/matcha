import React from 'react';
import Box from '@mui/material/Box';
import BackImg from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MainPage/backgroung-unsplash.jpg';
import Logo from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MenuSection/uletter.png';

const Swipe = () => {
	return (
		<div>
			<Box className="App__WebContainer__Website__Main">
				<img className="App__WebContainer__Website__Main__Background" src={BackImg} style={{ width:"100%"}}/>
				<Box className="App__WebContainer__Website__Main__OldOldCard">
					<Box className="App__WebContainer__Website__Main__OldOldCard__Profile">
						Je suis une tres vieille carte
					</Box>
				</Box>
				<Box className="App__WebContainer__Website__Main__OldCard">
					<Box className="App__WebContainer__Website__Main__OldCard__Profile">
						Je suis une vieille carte
					</Box>
				</Box>
				<Box className="App__WebContainer__Website__Main__PrimaryCard">
					<Box className="App__WebContainer__Website__Main__PrimaryCard__Profile">
						Je suis une carte
					</Box>
					<Box className="App__WebContainer__Website__Main__Response">
						<Box className="App__WebContainer__Website__Main__Response__No">
							Non
						</Box>
						<Box className="App__WebContainer__Website__Main__Response__Yes">
							Oui
						</Box>
					</Box>
				</Box>
				<Box className="App__WebContainer__Website__Main__NextCard">
					<img src={Logo} width={70} height={70} />
				</Box>
				<Box className="App__WebContainer__Website__Main__NextNextCard">
					<img src={Logo} width={50} height={50} opacity={0.2} />
				</Box>
			</Box>
		</div>
	);
}

export default Swipe;