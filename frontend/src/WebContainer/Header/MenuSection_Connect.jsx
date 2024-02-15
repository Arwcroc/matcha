// import React from 'react';
import * as React from 'react';
import Box from '@mui/material/Box';
import Logo from '../../Images/MenuSection/Urme-logo.png';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import FilterAltIcon from '@mui/icons-material/FilterAlt';
// import ForumOutlinedIcon from '@mui/icons-material/ForumOutlined';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';
// import Avatar from '@mui/material/Avatar';
import Tooltip from '@mui/material/Tooltip';
import Settings from '@mui/icons-material/Settings';
import Logout from '@mui/icons-material/Logout';
import Diversity1RoundedIcon from '@mui/icons-material/Diversity1Rounded';
import MailRoundedIcon from '@mui/icons-material/MailRounded';

import TextField from '@mui/material/TextField'
import FormControlLabel from '@mui/material/FormControlLabel'
import Checkbox from '@mui/material/Checkbox';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';

import Portrait from '../../Images/MainPage/joconde.jpg';
import GirlPort from '../../Images/MainPage/Girl_Pearl.jpg';
import Donna from '../../Images/MainPage/donna.jpg';
import Meuh from '../../Images/MainPage/meuhmeuh.jpg';
import Ermine from '../../Images/MainPage/Ermine.jpg';
import { CarouselProvider, Slider, Slide, ButtonBack, ButtonNext, DotGroup } from 'pure-react-carousel';
import 'pure-react-carousel/dist/react-carousel.es.css';
import { Typography } from '@mui/material';

const SlideBox = () => {
	return (
		<CarouselProvider
			naturalSlideWidth={100}
			naturalSlideHeight={100}
			totalSlides={5}
			className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Slider"
			infinite
		>
			<DotGroup className="App__WebContainer__Website__Main__PrimaryCard__Carousel__DotGroup" />
			<ButtonBack className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Back">{"<"}</ButtonBack>
			<Slider className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Pics">
				<Slide index={0}><div class="slide__image__container"><img class="slide__image__prout" src={Portrait} alt="Slide 1"/></div></Slide>
				<Slide index={1}><div class="slide__image__container"><img class="slide__image__prout" src={GirlPort} alt="Slide 2"/></div></Slide>
				<Slide index={2}><div class="slide__image__container"><img class="slide__image__prout" src={Meuh} alt="Slide 3"/></div></Slide>
				<Slide index={3}><div class="slide__image__container"><img class="slide__image__prout" src={Ermine} alt="Slide 4"/></div></Slide>
				<Slide index={4}><div class="slide__image__container"><img class="slide__image__prout" src={Donna} alt="Slide 5"/></div></Slide>
			</Slider>
			<Typography className="App__WebContainer__Header__ProfilePopUp__NameAge" fontSize={"24px"}>Sandrine, 14</Typography>
			<Typography className="App__WebContainer__Header__ProfilePopUp__Place" fontSize={"14px"}>Saint-Etienne</Typography>
			<ButtonNext className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Next">{">"}</ButtonNext>
		</CarouselProvider>
	);
}


const PopUp_Profile = () => {
	return (
		<>
			<Box className="App__WebContainer__Header__ProfilePopUp">
				<Box className="App__WebContainer__Header__ProfilePopUp__FirstBlock">
					<SlideBox />
				</Box>
				<Box className="App__WebContainer__Header__ProfilePopUp__Score">
					Score
				</Box>
				<Box>
					<Box className="App__WebContainer__Header__ProfilePopUp__Gender">
						Genre
					</Box>
					<Box className="App__WebContainer__Header__ProfilePopUp__SexualInterest">
						Sexual Interest
					</Box>
				</Box>
				<Box className="App__WebContainer__Header__ProfilePopUp__Bio">
					Biography
				</Box>
				<Box className="App__WebContainer__Header__ProfilePopUp__Tags">
					Interest
				</Box>
			</Box>
		</>
	);
}

const PopUp_Message = () => {
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

const PopUp_Setting = () => {
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

const PopUp_Matches = () => {
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

const MenuSection_Connect = () => {
	const [messageOpen, setMessageOpen] = React.useState(false);
	const [profileOpen, setProfileOpen] = React.useState(false);
	const [settingsOpen, setSettingsOpen] = React.useState(false);
	const [matchesOpen, setMatchesOpen] = React.useState(false);

	const [anchorEl, setAnchorEl] = React.useState(null);
	const listOpen = Boolean(anchorEl);


	const handleClick = (event) => {
	  setAnchorEl(event.currentTarget);
	};
	const closeList = () => {
	  setAnchorEl(null);
	};

	const openProfile = () => {
		closeList();
		setProfileOpen(true);
	};

	const openMessage = () => {
		closeList();
		setMessageOpen(true);
	};

	const openMatches = () => {
		closeList();
		setMatchesOpen(true);
	};

	const openSettings = () => {
		closeList();
		setSettingsOpen(true);
	};

	// const handleMessageClose = () => {
	// 	setMessageOpen(false);
	// };

	return (
		<>
			<Box className="App__WebContainer__Header__MenuSection">
				<Box className="App__WebContainer__Header__MenuSection__Logo">
					<img src={Logo} width={150} height={75} />
				</Box>
				<Box className="App__WebContainer__Header__MenuSection__RightSection">
					<Box className="App__WebContainer__Header__MenuSection__Login">
						<FilterAltIcon/>Filter
					</Box>
					<Box className="App__WebContainer__Header__MenuSection__MenuIcon">
						<Tooltip title="Account settings">
							<IconButton
								onClick={handleClick}
								size="small"
								sx={{ ml: 2 }}
								aria-controls={listOpen ? 'account-menu' : undefined}
								aria-haspopup="true"
								aria-expanded={listOpen ? 'true' : undefined}
							>
								<AccountCircleIcon className="App__WebContainer__Header__MenuSection__Avatar" fontSize='large'/>
							</IconButton>
						</Tooltip>
					</Box>
				</Box>
			</Box>
			<Menu
				anchorEl={anchorEl}
				id="account-menu"
				open={listOpen}
				onClose={closeList}
				onClick={closeList}
				PaperProps={{
				elevation: 0,
				sx: {
						overflow: 'visible',
						filter: 'drop-shadow(0px 8px 8px rgba(0,0,0,0.32))',
						mt: 1.5,
						'& .MuiAvatar-root': {
						width: 32,
						height: 32,
						ml: -0.5,
						mr: 1,
					},
					'&::before': {
						content: '""',
						display: 'block',
						position: 'absolute',
						top: 0,
						right: 25,
						width: 10,
						height: 10,
						bgcolor: 'background.paper',
						transform: 'translateY(-50%) rotate(45deg)',
						zIndex: 0,
						},
				},
				}}
				transformOrigin={{ horizontal: 'right', vertical: 'top' }}
				anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
			>
				<MenuItem onClick={openProfile}>
					<ListItemIcon>
          				<AccountCircleIcon fontSize='large' sx={{mr: 1}}/>
					</ListItemIcon>
					Profile
        		</MenuItem>
       			<MenuItem onClick={openMessage}>
					<ListItemIcon>
						<MailRoundedIcon fontSize='large' sx={{mr: 1}}/>
					</ListItemIcon>
					Messages
        		</MenuItem>
        		<Divider />
        		<MenuItem onClick={openMatches}>
          			<ListItemIcon>
           				 <Diversity1RoundedIcon fontSize="small" />
          			</ListItemIcon>
          			All matches
        		</MenuItem>
        		<MenuItem onClick={openSettings}>
          			<ListItemIcon>
            			<Settings fontSize="small" />
          			</ListItemIcon>
          			Settings
        		</MenuItem>
        		<MenuItem onClick={closeList}>
          			<ListItemIcon>
            			<Logout fontSize="small" />
          			</ListItemIcon>
          			Logout
        		</MenuItem>
      		</Menu>
			<Box>
				<Dialog open={profileOpen} onClose={() => setProfileOpen(false)}>
					<DialogContent className="App__WebContainer__Header__MenuSection__ProfilePopUp">
						<PopUp_Profile />
					</DialogContent>
				</Dialog>
				<Dialog open={messageOpen} onClose={() => setMessageOpen(false)}>
					<DialogContent className="App__WebContainer__Header__MenuSection__Dialog">
						<PopUp_Login />
					</DialogContent>
				</Dialog>
				<Dialog open={matchesOpen} onClose={() => setMatchesOpen(false)}>
					<DialogContent className="App__WebContainer__Header__MenuSection__Dialog">
						<PopUp_Login />
					</DialogContent>
				</Dialog>
				<Dialog open={settingsOpen} onClose={() => setSettingsOpen(false)}>
					<DialogContent className="App__WebContainer__Header__MenuSection__Dialog">
						<PopUp_Login />
					</DialogContent>
				</Dialog>
			</Box>
		</>
	)
}

export default MenuSection_Connect;