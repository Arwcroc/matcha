import React from 'react';
import Box from '@mui/material/Box';
import BackImg from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MainPage/backgroung-unsplash.jpg';
import Logo from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MenuSection/uletter.png';
import { CarouselProvider, Slider, Slide } from 'pure-react-carousel';
import 'pure-react-carousel/dist/react-carousel.es.css';

const SlideBox = () => {
	return (
	<CarouselProvider
		naturalSlideWidth={50}
		naturalSlideHeight={50}
		totalSlides={3}
	>
		<Slider>
			<Slide index={0}>
				<img src={BackImg} alt="Slide 1" />
			</Slide>
			<Slide index={1}>
				<img src={Logo} alt="Slide 2" />
			</Slide>
		 	<Slide index={2}>
				<img src={Logo} alt="Slide 3" />
			</Slide>
		</Slider>
	</CarouselProvider>
	);
}

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
						<SlideBox />
						Je suis pas la
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