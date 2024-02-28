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
import { styled } from '@mui/material/styles';
import Rating from '@mui/material/Rating';
import FavoriteIcon from '@mui/icons-material/Favorite';
import FavoriteBorderIcon from '@mui/icons-material/FavoriteBorder';

import Chip from '@mui/material/Chip';
import Stack from '@mui/material/Stack';

import Avatar from '@mui/material/Avatar';
import SearchIcon from '@mui/icons-material/Search';
import Badge from '@mui/material/Badge';
import ArrowCircleUpRoundedIcon from '@mui/icons-material/ArrowCircleUpRounded';

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
				<Slide index={0}><div className="slide__image__container"><img className="slide__image__prout" src={Portrait} alt="Slide 1"/></div></Slide>
				<Slide index={1}><div className="slide__image__container"><img className="slide__image__prout" src={GirlPort} alt="Slide 2"/></div></Slide>
				<Slide index={2}><div className="slide__image__container"><img className="slide__image__prout" src={Meuh} alt="Slide 3"/></div></Slide>
				<Slide index={3}><div className="slide__image__container"><img className="slide__image__prout" src={Ermine} alt="Slide 4"/></div></Slide>
				<Slide index={4}><div className="slide__image__container"><img className="slide__image__prout" src={Donna} alt="Slide 5"/></div></Slide>
			</Slider>
			<Typography className="App__WebContainer__Header__ProfilePopUp__NameAge" fontSize={"24px"}>Sandrine, 14</Typography>
			<Typography className="App__WebContainer__Header__ProfilePopUp__Place" fontSize={"14px"}>Saint-Etienne</Typography>
			<ButtonNext className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Next">{">"}</ButtonNext>
		</CarouselProvider>
	);
}

const StyledRating = styled(Rating)({
	'& .MuiRating-iconFilled': {
	  color: '#ff6d75',
	},
	'& .MuiRating-iconHover': {
	  color: '#ff3d47',
	},
});

const PopUp_Profile = () => {
	return (
		<>
			<Box className="App__WebContainer__Header__ProfilePopUp">
				<Box className="App__WebContainer__Header__ProfilePopUp__FirstBlock">
					<SlideBox />
				</Box>
				<Box className="App__WebContainer__Header__ProfilePopUp__Score">
					{/* <Rating name="OrgasmScore" value={4} readOnly /> */}
					<StyledRating
						name="customized-color"
						defaultValue={4.6}
						getLabelText={(value) => `${value} Heart${value !== 1 ? 's' : ''}`}
						precision={0.1}
						icon={<FavoriteIcon fontSize="inherit" />}
						emptyIcon={<FavoriteBorderIcon fontSize="inherit" />}
						readOnly
					/>
				</Box>
				<Box className="App__WebContainer__Header__ProfilePopUp__SexInfo">
					<Box className="App__WebContainer__Header__ProfilePopUp__Gender">
						Genre
						<Stack direction="row" spacing={1}>
							<Chip label="Homme" />
						</Stack>
					</Box>
					<Box className="App__WebContainer__Header__ProfilePopUp__SexualInterest">
						Sexual Interest
						<Stack direction="row" spacing={1}>
							<Chip label="Homme" />
							<Chip label="Femme" variant="outlined" />
							<Chip label="Autre" />
						</Stack>
					</Box>
				</Box>
				<Box className="App__WebContainer__Header__ProfilePopUp__Bio">
					Biography
					<Box className="App__WebContainer__Header__ProfilePopUp__BioText">
						Voyageur passionné, rêveur épicurien. 
						Artiste de la vie cherchant complicité et partage. 
						Sourire contagieux, esprit curieux. Prêt pour l'aventure.
					</Box>
				</Box>
				<Box className="App__WebContainer__Header__ProfilePopUp__TagsTitle">
					Interest
				</Box>
				<Box className="App__WebContainer__Header__ProfilePopUp__Tags">
					<Chip className='App__WebContainer__Header__ProfilePopUp__Tags_Tag' label="#3dPrint" />
					<Chip className='App__WebContainer__Header__ProfilePopUp__Tags_Tag' label="#Fist"/>
					<Chip className='App__WebContainer__Header__ProfilePopUp__Tags_Tag' label="#Choucroute" />
					<Chip className='App__WebContainer__Header__ProfilePopUp__Tags_Tag' label="#GrosseVoiture" />
					<Chip className='App__WebContainer__Header__ProfilePopUp__Tags_Tag' label="#VroumVroum" />
					<Chip className='App__WebContainer__Header__ProfilePopUp__Tags_Tag' label="#France" />
				</Box>
			</Box>
		</>
	);
}

const PopUp_Message = () => {
	return (
		<>
			<Box className="App__WebContainer__Header__MessagePopUp">
				<Box className="App__WebContainer__Header__MessagePopUp__GuysSection">
					<Box className="App__WebContainer__Header__MessagePopUp__GuysSection__Title">Chat</Box>
					<Divider/>
					<Box className="App__WebContainer__Header__MessagePopUp__GuysSection__SearchBox">
						<SearchIcon className="App__WebContainer__Header__MessagePopUp__GuysSection__SearchIcon"/>
						<input className="App__WebContainer__Header__MessagePopUp__GuysSection__Search" placeholder="Search..." type="text"/>
					</Box>
					<Box className="App__WebContainer__Header__MessagePopUp__GuysSection__UserBox">
						<Badge color="success" overlap="circular" badgeContent=" " variant="dot">
							<Avatar className="App__WebContainer__Header__MessagePopUp__GuysSection__UserPics" src={Meuh}>
								A
							</Avatar>
						</Badge>		
						<p className="App__WebContainer__Header__MessagePopUp__GuysSection__UserInfo">Anne</p>
					</Box>
					<Box className="App__WebContainer__Header__MessagePopUp__GuysSection__UserBox">
						<Badge color="error" overlap="circular" badgeContent=" " variant="dot">
							<Avatar className="App__WebContainer__Header__MessagePopUp__GuysSection__UserPics" src={Donna}>
								A
							</Avatar>
						</Badge>		
						<p className="App__WebContainer__Header__MessagePopUp__GuysSection__UserInfo">Sandrine</p>
					</Box>
				</Box>
				<Box className="App__WebContainer__Header__MessagePopUp__MessageSection">
					<Box className="App__WebContainer__Header__MessagePopUp__MessageSection__UserInfo">Anne</Box>
					<Divider/>
					<Box className="App__WebContainer__Header__MessagePopUp__MessageSection__BlockMessage">
						<Box className="App__WebContainer__Header__MessagePopUp__MessageSection__MsgIn">message in and its for the fame you know, like when taylor swift take the biggest airplane to go to take the break, the oui oui baguette </Box>
						<Box className="App__WebContainer__Header__MessagePopUp__MessageSection__MsgOut">message out</Box>
					</Box>
					<Box className="App__WebContainer__Header__MessagePopUp__MessageSection__TypeZone">
						<textarea onkeyup="textAreaAdjust(this)" className="App__WebContainer__Header__MessagePopUp__MessageSection__InputType"/>
						<ArrowCircleUpRoundedIcon className="App__WebContainer__Header__MessagePopUp__MessageSection__SendIcon"/>
					</Box>
				</Box>
			</Box>
		</>
	);
}

const PopUp_Setting = () => {
	return (
		<>
			<Box className="App__WebContainer__Header__SettingPopUp">
				<Box className="App__WebContainer__Header__SettingPopUp_Title">
					Settings
				</Box>
				<Box>
					Name
				</Box>
				<Box>
					Last Name
				</Box>
				<Box>
					Gender
				</Box>
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
					<Box className="App__WebContainer__Header__MenuSection__MessagePopUp">
						<PopUp_Message />
					</Box>
				</Dialog>
				<Dialog open={settingsOpen} onClose={() => setSettingsOpen(false)} sx={{'.MuiDialogContent-root': {padding: 0, borderRadius: "8px"},}} >
					<DialogContent className="App__WebContainer__Header__MenuSection__SettingPopUp">
						<PopUp_Setting />
					</DialogContent>
				</Dialog>
				<Dialog open={matchesOpen} onClose={() => setMatchesOpen(false)}>
					<DialogContent className="App__WebContainer__Header__MenuSection__Dialog">
						<PopUp_Login />
					</DialogContent>
				</Dialog>
			</Box>
		</>
	)
}

export default MenuSection_Connect;