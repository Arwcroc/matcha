import React from 'react';
import Box from '@mui/material/Box';
import Logo from '../../Images/MenuSection/Urme-logo.png';
import MenuIcon from '@mui/icons-material/Menu';

const MenuSection_Connect = () => {
	return (
		<Box className="App__WebContainer__Header__MenuSection">
			<Box className="App__WebContainer__Header__MenuSection__Logo">
				<img src={Logo} width={150} height={75} />
			</Box>
			<Box className="App__WebContainer__Header__MenuSection__RightSection">
				<Box className="App__WebContainer__Header__MenuSection__Login">
					Account
				</Box>
				<Box className="App__WebContainer__Header__MenuSection__MenuIcon">
					<MenuIcon/>
				</Box>
			</Box>
		</Box>
	)
}

export default MenuSection_Connect;