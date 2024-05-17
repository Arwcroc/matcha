// import React from 'react';
import * as React from 'react';
import { useState } from 'react';
import Box from '@mui/material/Box';
import Logo from '../../Images/MenuSection/Urme-logo.png';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import FilterAltIcon from '@mui/icons-material/FilterAlt';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import Settings from '@mui/icons-material/Settings';
import Logout from '@mui/icons-material/Logout';
import Diversity1RoundedIcon from '@mui/icons-material/Diversity1Rounded';
import MailRoundedIcon from '@mui/icons-material/MailRounded';

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

import EditIcon from '@mui/icons-material/Edit';
import VisibilityIcon from '@mui/icons-material/Visibility';
import ReportOutlinedIcon from '@mui/icons-material/ReportOutlined';
import RemoveShoppingCartOutlinedIcon from '@mui/icons-material/RemoveShoppingCartOutlined';
import BlockOutlinedIcon from '@mui/icons-material/BlockOutlined';

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
	const [messages, setMessages] = useState({});
    const [inputText, setInputText] = useState('');
	const [selectedUser, setSelectedUser] = useState(null);
	const [selectedUserName, setSelectedUserName] = useState('');

    const handleInputChange = (e) => {
        setInputText(e.target.value);
    };

    const sendMessage = () => {
        if (inputText.trim() !== '') {
            const newMessage = {
                text: inputText,
                sender: 'me',
            };
			const userMessages = messages[selectedUserName] ? messages[selectedUserName] : []
			const newMessages = {...messages}
			newMessages[selectedUserName] = [...userMessages, newMessage]
            setMessages(newMessages);
            setInputText('');
        }
    };

    const handleKeyPress = (e) => {
		if (e.key === 'Enter') {
			e.preventDefault();
            sendMessage();
        }
    };

	const handleUserClick = (userName) => {
        setSelectedUser(userName);
        setSelectedUserName(userName);
    };

	const displayedMessages = messages[selectedUserName] ? messages[selectedUserName] : []

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
					<Box className="App__WebContainer__Header__MessagePopUp__GuysSection__UserBox" onClick={() => handleUserClick('Anne')}>
						<Badge color="success" overlap="circular" badgeContent=" " variant="dot">
							<Avatar className="App__WebContainer__Header__MessagePopUp__GuysSection__UserPics" src={Meuh}>
								A
							</Avatar>
						</Badge>		
						<p className="App__WebContainer__Header__MessagePopUp__GuysSection__UserInfo">Anne</p>
					</Box>
					<Box className="App__WebContainer__Header__MessagePopUp__GuysSection__UserBox" onClick={() => handleUserClick('Sandrine')}>
						<Badge color="error" overlap="circular" badgeContent=" " variant="dot">
							<Avatar className="App__WebContainer__Header__MessagePopUp__GuysSection__UserPics" src={Donna}>
								S
							</Avatar>
						</Badge>		
						<p className="App__WebContainer__Header__MessagePopUp__GuysSection__UserInfo">Sandrine</p>
					</Box>
				</Box>
				<Box className="App__WebContainer__Header__MessagePopUp__MessageSection">
					<Box className="App__WebContainer__Header__MessagePopUp__MessageSection__UserInfo">{selectedUserName !== '' ? selectedUserName : "Profil"}</Box>
					<Divider/>
					<Box className="App__WebContainer__Header__MessagePopUp__MessageSection__BlockMessage">
						{displayedMessages.map((message, index) => {
							if (selectedUser === null || message.sender === selectedUser || message.sender === 'me') {
								return (
									<Box
										key={index}
										className={`App__WebContainer__Header__MessagePopUp__MessageSection__Msg${message.sender === 'me' ? 'Out' : 'In'}`}
									>
										{message.text}
									</Box>
								);
							} else {
								return null;
							}
						})}
					</Box>
					<Box className="App__WebContainer__Header__MessagePopUp__MessageSection__TypeZone">
						<textarea 
							onkeyup="textAreaAdjust(this)"
							className="App__WebContainer__Header__MessagePopUp__MessageSection__InputType"
							value={inputText}
							onChange={handleInputChange}
                        	onKeyPress={handleKeyPress}
							/>
						<ArrowCircleUpRoundedIcon className="App__WebContainer__Header__MessagePopUp__MessageSection__SendIcon" onClick={sendMessage}/>
					</Box>
				</Box>
			</Box>
		</>
	);
}

const valueMapping = [
	{ step: 0, label: 'Autour de moi' },
	{ step: 1, label: '10km' },
	{ step: 2, label: '15km' },
	{ step: 3, label: '20km' },
	{ step: 4, label: '30km' },
	{ step: 5, label: '50km' },
	{ step: 6, label: '75km' },
	{ step: 7, label: '100km' },
	{ step: 8, label: '150km' },
	{ step: 9, label: '200km' },
	{ step: 10, label: '250km' },
	{ step: 11, label: '300km' },
	{ step: 12, label: '400km' },
	{ step: 13, label: "J'ai faim" },
];

const valueAgeMapping = [
	{ step: 0, label: 'Pile pile' },
	{ step: 1, label: "1 an d'écart" },
	{ step: 2, label: "2 ans d'écart" },
	{ step: 3, label: "3 ans d'écart" },
	{ step: 4, label: "4 ans d'écart" },
	{ step: 5, label: "5 ans d'écart" },
	{ step: 6, label: "10 ans d'écart" },
	{ step: 7, label: "15 ans d'écart" },
	{ step: 8, label: "20 ans d'écart" },
	{ step: 9, label: "25 ans d'écart" },
	{ step: 10, label: "30 ans d'écart" },
	{ step: 11, label: "35 ans d'écart" },
	{ step: 12, label: "40 ans d'écart" },
	{ step: 13, label: "J'ai faim" },
];
  
const LimitSlider = (props) => {
	const {min, max, mapSelector, name} = props;
	const [value, setValue] = useState(0);

	const handleChange = (event) => {
		const newValue = parseInt(event.target.value, 10);
		setValue(newValue);
	};

	return (
		<div className="SettingPopUp__LimitSlider">
			<p>{name} {mapSelector[value]?.label || "cpt"}</p>
			<input
				type="range"
				min={min}
				max={max}
				step="1"
				value={value}
				onChange={handleChange}
				className="SettingPopUp__LimitSlider__Valor"
			/>
		</div>
	);
};

const Switch = () => {
	const [isChecked, setIsChecked] = useState(false);
  
	const toggleSwitch = () => {
		setIsChecked(!isChecked);
	};
  
	return (
		<div className={`switch ${isChecked ? 'on' : 'off'}`} onClick={toggleSwitch}>
			<div className={`slider ${isChecked ? 'on' : 'off'}`}></div>
		</div>
	);
};

const getPosition = (state) => {
	switch (state) {
		case 'Homme':
		case 'Hétéro':
			return 0;
		case 'Autre':
		case 'Bi':
			return 90;
		case 'Femme':
		case 'Lesb/Gay':
			return 180;
		default:
			return 0;
	}
};


const ThreeStateToggleSwitch = (props) => {
	const { prop1, prop2, prop3 } = props;
	const [toggleState, setToggleState] = useState(prop2);

	const handleToggle = (newState) => {
		setToggleState(newState);
	};
  
	return (
		<div className="SettingPopUp__TriToggleSwitch">
			<div className="SettingPopUp__TriToggleSwitch__Slider" style={{ transform: `translateX(${getPosition(toggleState)}px)` }}>
				{toggleState}
			</div>
			<div className="SettingPopUp__TriToggleSwitch__Slider__Option First" onClick={() => handleToggle(prop1)}>{prop1}</div>
			<div className="SettingPopUp__TriToggleSwitch__Slider__Option Second" onClick={() => handleToggle(prop2)}>{prop2}</div>
			<div className="SettingPopUp__TriToggleSwitch__Slider__Option Third" onClick={() => handleToggle(prop3)}>{prop3}</div>
		</div>
	);
};

const AgeSelect = ({ onChange, selectedAge }) => {
	const ageOptions = [];
	for (let i = 10; i <= 100; i++) {
		ageOptions.push(
		<option key={i} value={i}>
			{i}
		</option>
		);
	}
	
	return (
		<select className="App__WebContainer__Header__SettingPopUp__SettingName__Select" onChange={(e) => onChange(e.target.value)} value={selectedAge}>
			{ageOptions}
		</select>
	);
};

const RateButtons = () => {
	const [selectedRate, setSelectedRate] = useState(null);
	
	const handleRateSelection = (rate) => {
		setSelectedRate(rate === selectedRate ? null : rate);
	};
	
	return (
		<div className="App__WebContainer__Header__SettingPopUp__SettingName__AllRate">
		{[4, 3, 2, 1, 0].map((rate) => (
			<div
			key={rate}
			className={`App__WebContainer__Header__SettingPopUp__SettingName__RateContainer ${rate === selectedRate ? 'selected' : 'unselected'}`}
			onClick={() => handleRateSelection(rate)}
			>
			<p className="App__WebContainer__Header__SettingPopUp__SettingName__RateSelection">{`${rate} et +`}</p>
			</div>
		))}
		</div>
	);
};

const PopUp_Setting = () => {
	const [selectedAge, setSelectedAge] = useState(18);

	const handleAgeChange = (age) => {
	  setSelectedAge(age);
	};

	return (
		<>
			<Box className="App__WebContainer__Header__SettingPopUp">
				<Box className="App__WebContainer__Header__SettingPopUp__Title">
					Settings
				</Box>
				<Box className="App__WebContainer__Header__SettingPopUp__BoxTitle">
					Personnal Information
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName__SimpleChange">
						Name
						<EditIcon className="App__WebContainer__Header__SettingPopUp__SettingName__Icon" style={{ fontSize: '15px' }}/>
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">OMG your name offends me, you can change it by clicking on the icon just right here</p>
						<div className="App__WebContainer__Header__SettingPopUp__SettingName__DivBoxName">
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__NameBox">Name</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__NameBox">Last Name</Box>
						</div>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName__SimpleChange">
						Email
						<EditIcon className="App__WebContainer__Header__SettingPopUp__SettingName__Icon" style={{ fontSize: '15px' }}/>
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Change your caramail</p>
						<div className="App__WebContainer__Header__SettingPopUp__SettingName__DivBoxName">
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__NameBox">sac@merde.com</Box>
						</div>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName__SimpleChange">
						Age
						<EditIcon className="App__WebContainer__Header__SettingPopUp__SettingName__Icon" style={{ fontSize: '15px' }}/>
      					<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Selected the age you want to show</p>
						<AgeSelect onChange={handleAgeChange} selectedAge={selectedAge} />
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName TriSwitch">
						Gender
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Choose a gender, yes you have the choice...</p>
						<ThreeStateToggleSwitch prop1='Homme' prop2='Autre' prop3='Femme'/>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName TriSwitch">
						Sexual Interest
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Only steers and queers in Texas</p>
						<ThreeStateToggleSwitch prop1='Hétéro' prop2='Bi' prop3='Lesb/Gay'/>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName__FlagBox">
						Personal Tag
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Not here to kinkshame, dirtyboy</p>
						<Box className="App__WebContainer__Header__SettingPopUp__SettingName__AllTag">
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__Tag">Naruto</p>
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagRemover">x</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__Tag">Riz</p>
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagRemover">x</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagAdd">+</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagAdd">+</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagAdd">+</p>
							</Box>
						</Box>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName__SimpleChange">
						Biography
						<EditIcon className="App__WebContainer__Header__SettingPopUp__SettingName__Icon" style={{ fontSize: '15px' }}/>
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Talk about yourself, a usual</p>
						<Box className="App__WebContainer__Header__SettingPopUp__SettingName__NameBox">J'apprécie les fruits en sirop</Box>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName__Picture">
						Picture
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Try to choose one where you wear clothes...</p>
						<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer">
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__PicsBox">
								<img className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__Pics" src={Portrait} alt={"profile1"}/>
								<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__Close">x</Box>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__PicsBox">
								<img className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__Pics" src={Meuh} alt={"profile2"}/>
								<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__Close">x</Box>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__PicsBox">
								<img className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__Pics" src={Donna} alt={"profile3"}/>
								<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__Close">x</Box>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__PicsBox">
								<img className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__Pics" src={Ermine} alt={"profile4"}/>
								<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__Close">x</Box>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__PicsBox">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__PictureContainer__Nothing">+</p>
							</Box>
						</Box>
					</Box>
				</Box>
				<Box className="App__WebContainer__Header__SettingPopUp__BoxTitle">
					Services
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName Switch">
						<Box>
							Location
							<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Make your location great again</p>
						</Box>
						<Switch />
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName Switch">
						<Box>
							Active Wi-Fi
							<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Enable your Wi-Fi</p>
						</Box>
						<Switch />
					</Box>
				</Box>
				<Box className="App__WebContainer__Header__SettingPopUp__BoxTitle">
					Filter
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName TriSwitch">
						Age Gap
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Select what you want, it isn't me who go to jail !</p>
						<LimitSlider min="0" max="13" mapSelector={valueAgeMapping} name="Age : "/>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName TriSwitch">
						Geographical Limit
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">All roads lead to bum...</p>
						<LimitSlider min="0" max="13" mapSelector={valueMapping} name="Distance : "/>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName__FlagBox">
						Interest Flag
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">Not here to kinkshame, dirtyboy</p>
						<Box className="App__WebContainer__Header__SettingPopUp__SettingName__AllTag">
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__Tag">Naruto</p>
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagRemover">x</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__Tag">Riz</p>
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagRemover">x</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagAdd">+</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagAdd">+</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__TagContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__TagAdd">+</p>
							</Box>
						</Box>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__SettingName__FlagBox">
						Fame Rating
						<p className="App__WebContainer__Header__SettingPopUp__SettingName__Text">The level of hornyness you want to encounter</p>
						{/* <Box className="App__WebContainer__Header__SettingPopUp__SettingName__AllRate">
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__RateContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__RateSelection">4 et +</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__RateContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__RateSelection">3 et +</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__RateContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__RateSelection">2 et +</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__RateContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__RateSelection">1 et +</p>
							</Box>
							<Box className="App__WebContainer__Header__SettingPopUp__SettingName__RateContainer">
								<p className="App__WebContainer__Header__SettingPopUp__SettingName__RateSelection">0 et +</p>
							</Box>
						</Box> */}
						<RateButtons/>
					</Box>
				</Box>
			</Box>
		</>
	);
}

const LikedPerson = (props) => {
	const {altProfile, picsName} = props;

	return (
		<Box className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsBox" alt={altProfile}>
			<img className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__Pics" src={picsName} alt={altProfile}/>
			<ReportOutlinedIcon className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsTag Report" fontSize='small'/>
			<RemoveShoppingCartOutlinedIcon className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsTag Remove" fontSize='small'/>
			<BlockOutlinedIcon className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsTag Block" fontSize='small'/>
		</Box>
	);
}

const PopUp_Matches = () => {
	return (
		<>
			<Box className="App__WebContainer__Header__SettingPopUp__Interact">
				<Box className="App__WebContainer__Header__SettingPopUp__Interact__Title">
					Intereacts
				</Box>
				<Divider className="App__WebContainer__Header__SettingPopUp__Interact__Divider" sx={{ borderBottomWidth: 2 }} style={{ background: 'black' }}/>
				<p className="App__WebContainer__Header__SettingPopUp__Interact__SmallTitle">People I liked</p>
				<Box className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer">
					<LikedPerson altProfile="profile1" picsName={Portrait}/>
					<LikedPerson altProfile="profile2" picsName={Meuh}/>
					<LikedPerson altProfile="profile3" picsName={Donna}/>
					<LikedPerson altProfile="profile4" picsName={Ermine}/>
					<LikedPerson altProfile="profile1" picsName={Portrait}/>
				</Box>
				<Divider className="App__WebContainer__Header__SettingPopUp__Interact__Divider" sx={{ borderBottomWidth: 2 }} style={{ background: 'black' }}/>
				<p className="App__WebContainer__Header__SettingPopUp__Interact__SmallTitle">People interact with my profile</p>
				<Box className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer">
					<Box className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsBox" alt={"profile1"}>
						<img className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__Pics viewed" src={Portrait} alt={"profile1"}/>
						<VisibilityIcon className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsTag view" fontSize='small'/>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsBox" alt={"profile2"}>
						<img className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__Pics liked" src={Meuh} alt={"profile2"}/>
						<FavoriteIcon className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsTag like" fontSize='small'/>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsBox" alt={"profile3"}>
						<img className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__Pics liked" src={Donna} alt={"profile3"}/>
						<FavoriteIcon className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsTag like" fontSize='small'/>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsBox" alt={"profile4"}>
						<img className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__Pics viewed" src={Ermine} alt={"profile4"}/>
						<VisibilityIcon className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsTag view" fontSize='small'/>
					</Box>
					<Box className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsBox" alt={"profile1"}>
						<img className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__Pics viewed" src={Portrait} alt={"profile1"}/>
						<VisibilityIcon className="App__WebContainer__Header__SettingPopUp__Interact__PicsContainer__PicsTag view" fontSize='small'/>
					</Box>
				</Box>
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

	return (
		<>
			<Box className="App__WebContainer__Header__MenuSection">
				<Box className="App__WebContainer__Header__MenuSection__Logo">
					<img src={Logo} width={150} height={75} alt="logo"/>
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
          			Interact
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
				<Dialog open={matchesOpen} onClose={() => setMatchesOpen(false)} sx={{'.MuiDialogContent-root': {padding: 0, borderRadius: "8px"},}}>
					<DialogContent className="App__WebContainer__Header__MenuSection__Dialog">
						<PopUp_Matches />
					</DialogContent>
				</Dialog>
			</Box>
		</>
	)
}

export default MenuSection_Connect;