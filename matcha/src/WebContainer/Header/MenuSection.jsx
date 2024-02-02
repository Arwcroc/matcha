import React from 'react';
import Box from '@mui/material/Box';
import Logo from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MenuSection/Urme-logo.png';

import TextField from '@mui/material/TextField'
import FormControlLabel from '@mui/material/FormControlLabel'
import Checkbox from '@mui/material/Checkbox';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';

const PopUp_Login = () => {
	return (
		<>
			<Box className="App__WebContainer__Header__PopUp">
				<Box className="App__WebContainer__Header__PopUp__Name">
					Login
				</Box>
				<Box className="App__WebContainer__Header__PopUp__PassMail">
					<TextField id="email" label="Email" variant="standard" className="App__WebContainer__Header__PopUp__Email"/>
					<TextField id="password" label="Password" variant="standard" className="App__WebContainer__Header__PopUp__Password"/>
				</Box>
				<Box className="App__WebContainer__Header__PopUp__UsualAsk">
					<FormControlLabel control={<Checkbox />} label="Remember me" />
					<Box className="App__WebContainer__Header__PopUp__UsualAsk__Forgot">Forgot Password ?</Box>
				</Box>
				<Box className="App__WebContainer__Header__PopUp__Validate">Login</Box>
				<Box className="App__WebContainer__Header__PopUp__CreateAccount">Click here to create an account !</Box>
			</Box>
		</>
	);
}

const MenuSection = () => {
	const [open, setOpen] = React.useState(false);

	const handleClickOpen = () => {
		setOpen(true);
	};

	const handleClose = () => {
		setOpen(false);
	};
	return (
		<Box className="App__WebContainer__Header__MenuSection">
			<Box className="App__WebContainer__Header__MenuSection__Logo">
				<img src={Logo} width={150} height={75} />
			</Box>
			<Box className="App__WebContainer__Header__MenuSection__Login" onClick={handleClickOpen}>
				Login
			</Box>
			<Dialog open={open} onClose={handleClose}>
				<DialogContent className="App__WebContainer__Header__MenuSection__Dialog">
					<PopUp_Login />
				</DialogContent>
			</Dialog>
		</Box>
	)
}

export default MenuSection;