import React from 'react';
import Box from '@mui/material/Box';
import Logo from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MenuSection/Urme-logo.png';

const MenuSection = () => {
	return (
		<Box className="App__WebContainer__Header__MenuSection">
			<Box className="App__WebContainer__Header__MenuSection__Logo">
				<img src={Logo} width={150} height={75} />
			</Box>
			<Box className="App__WebContainer__Header__MenuSection__Login">
				Login
			</Box>
		</Box>
	)
}

export default MenuSection;