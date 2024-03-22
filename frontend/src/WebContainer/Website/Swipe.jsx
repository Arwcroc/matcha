import React from 'react';
import { useState } from 'react';
import Box from '@mui/material/Box';
import BackImg from '../../Images/MainPage/backgroung-unsplash.jpg';
import Portrait from '../../Images/MainPage/joconde.jpg';
import GirlPort from '../../Images/MainPage/Girl_Pearl.jpg';
import Donna from '../../Images/MainPage/donna.jpg';
import Meuh from '../../Images/MainPage/meuhmeuh.jpg';
import Ermine from '../../Images/MainPage/Ermine.jpg';
import { CarouselProvider, Slider, Slide, ButtonBack, ButtonNext, DotGroup } from 'pure-react-carousel';
import 'pure-react-carousel/dist/react-carousel.es.css';

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
			<ButtonNext className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Next">{">"}</ButtonNext>
		</CarouselProvider>
	);
}

const Primary = () => {
	const [currentIndex, setCurrentIndex] = useState(0);

	const handleLike = () => {
	  setCurrentIndex(currentIndex + 1);
	};
  
	const handleDislike = () => {
	  setCurrentIndex(currentIndex + 1);
	};

	return (
		<Box className="App__WebContainer__Website__Main__PrimaryCard">
			<Box className="App__WebContainer__Website__Main__PrimaryCard__Profile">
				<SlideBox />
			</Box>
			<Box className="App__WebContainer__Website__Main__Response">
				<Box className="App__WebContainer__Website__Main__Response__No" onClick={handleDislike}>
					Non
				</Box>
				<Box className="App__WebContainer__Website__Main__Response__Yes" onClick={handleLike}>
					Oui
				</Box>
			</Box>
		</Box>
	);
}

const Swipe = () => {
	return (
		<div>
			<Box className="App__WebContainer__Website__Main">
				<img className="App__WebContainer__Website__Main__Background" src={BackImg} style={{ width:"100%"}}/>
					<Box className="App__WebContainer__Website__Main__OldCard">
						<img src={Meuh} className="App__WebContainer__Website__Main__OldCard__Profile"/>
					</Box>
					<Primary />
					<Box className="App__WebContainer__Website__Main__NextCard">
						<img src={Donna} className="App__WebContainer__Website__Main__OldCard__Profile"/>
					</Box>
			</Box>
		</div>
	);
}

export default Swipe;