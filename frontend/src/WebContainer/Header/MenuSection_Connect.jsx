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
import Avatar from '@mui/material/Avatar';
import Tooltip from '@mui/material/Tooltip';
import Settings from '@mui/icons-material/Settings';
import Logout from '@mui/icons-material/Logout';
import Diversity1RoundedIcon from '@mui/icons-material/Diversity1Rounded';
import MailRoundedIcon from '@mui/icons-material/MailRounded';

const MenuSection_Connect = () => {
	const [anchorEl, setAnchorEl] = React.useState(null);
	const open = Boolean(anchorEl);
	const handleClick = (event) => {
	  setAnchorEl(event.currentTarget);
	};
	const handleClose = () => {
	  setAnchorEl(null);
	};
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
								aria-controls={open ? 'account-menu' : undefined}
								aria-haspopup="true"
								aria-expanded={open ? 'true' : undefined}
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
				open={open}
				onClose={handleClose}
				onClick={handleClose}
				PaperProps={{
				elevation: 0,
				sx: {
						overflow: 'visible',
						filter: 'drop-shadow(0px 2px 8px rgba(0,0,0,0.32))',
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
				<MenuItem onClick={handleClose}>
          			<Avatar /> Profile
        		</MenuItem>
       			<MenuItem onClick={handleClose}>
					<ListItemIcon>
        				<MailRoundedIcon fontSize='medium'/>
					</ListItemIcon>
					Messages
        		</MenuItem>
        		<Divider />
        		<MenuItem onClick={handleClose}>
          			<ListItemIcon>
           				 <Diversity1RoundedIcon fontSize="small" />
          			</ListItemIcon>
          			All matches
        		</MenuItem>
        		<MenuItem onClick={handleClose}>
          			<ListItemIcon>
            			<Settings fontSize="small" />
          			</ListItemIcon>
          			Settings
        		</MenuItem>
        		<MenuItem onClick={handleClose}>
          			<ListItemIcon>
            			<Logout fontSize="small" />
          			</ListItemIcon>
          			Logout
        		</MenuItem>
      		</Menu>
		</>
	)
}

export default MenuSection_Connect;