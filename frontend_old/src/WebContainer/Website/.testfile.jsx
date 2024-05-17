import React from 'react';
import Box from '@mui/material/Box';
import BackImg from '../../Images/MainPage/backgroung-unsplash.jpg';
import Portrait from '../../Images/MainPage/joconde.jpg';
import GirlPort from '../../Images/MainPage/Girl_Pearl.jpg';
import Donna from '../../Images/MainPage/donna.jpg';
import Meuh from '../../Images/MainPage/meuhmeuh.jpg';
import Ermine from '../../Images/MainPage/Ermine.jpg';
import { CarouselProvider, Slider, Slide, ButtonBack, ButtonNext, DotGroup } from 'pure-react-carousel';
import 'pure-react-carousel/dist/react-carousel.es.css';

import { Swiper, SwiperSlide } from 'swiper/react';
import { EffectCoverflow, Pagination } from 'swiper/modules';
import 'swiper/swiper-bundle.css';

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

const Swipe = () => {
	return (
		<div>
			<Box className="App__WebContainer__Website__Main">
				<img className="App__WebContainer__Website__Main__Background" src={BackImg} style={{ width:"100%"}}/>
				<Swiper
					effect={'coverflow'}
					grabCursor={true}
					// spaceBetween={50}
					slidesPerView={3}
					centeredSlides={true}
					// allowSlideNext={false}
					allowSlidePrev={false}
					coverflowEffect={{
						rotate: 50,
						stretch: 0,
						depth: 100,
						modifier: 1,
						slideShadows: true,
					}}
					modules={[EffectCoverflow]}
					className="mySwiper"
					onSlideChange={(swiper) => {
						const slides = swiper.slides;
						slides.forEach((slide) => {
						  slide.classList.remove('active-slide');
						});
						slides[swiper.activeIndex].classList.add('active-slide');
					  }}
				>
				<SwiperSlide>
					<Box className="App__WebContainer__Website__Main__OldOldCard">
						<img src={Portrait} className="App__WebContainer__Website__Main__OldOldCard__Profile"/>
					</Box>
				</SwiperSlide>
				<SwiperSlide>
					<Box className="App__WebContainer__Website__Main__OldCard">
						<img src={Portrait} className="App__WebContainer__Website__Main__OldCard__Profile"/>
					</Box>
				</SwiperSlide>
				<SwiperSlide>
					<Box className="App__WebContainer__Website__Main__PrimaryCard">
						<Box className="App__WebContainer__Website__Main__PrimaryCard__Profile">
							<SlideBox />
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
				</SwiperSlide>
				<SwiperSlide>
					<Box className="App__WebContainer__Website__Main__NextCard">
						<img src={Portrait} />
					</Box>
				</SwiperSlide>
				<SwiperSlide>
					<Box className="App__WebContainer__Website__Main__NextNextCard">
						<img src={Portrait} />
					</Box>
				</SwiperSlide>
				</Swiper>
			</Box>
		</div>
	);
}

export default Swipe;