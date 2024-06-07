import {
  Box,
  AppBar as MuiAppBar,
  Toolbar,
  Typography,
  Link,
  Stack,
} from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import { HelpFilledIcon, PmmRoundedIcon } from 'icons';
import { Breadcrumbs } from 'components/breadcrumbs';
import { PMM_HOME_URL } from 'constants';
import { Messages } from './AppBar.messages';

export const AppBar = () => (
  <MuiAppBar position="static" color="primary">
    <Toolbar>
      <Link
        href={PMM_HOME_URL}
        color="inherit"
        underline="hover"
        sx={(theme) => ({
          marginRight: theme.spacing(2),
        })}
        data-testid="appbar-pmm-link"
      >
        <Stack gap={1} direction="row" alignItems="center">
          <PmmRoundedIcon sx={{ height: '40px', width: 'auto' }} />
          <Typography>{Messages.title}</Typography>
        </Stack>
      </Link>
      <Breadcrumbs />
      <Box sx={{ ml: 'auto' }}>
        <Link
          to="/"
          component={RouterLink}
          color="inherit"
          underline="hover"
          data-testid="appbar-support-link"
        >
          <Stack gap={1} direction="row" alignItems="center">
            <HelpFilledIcon />
            <Typography>{Messages.support}</Typography>
          </Stack>
        </Link>
      </Box>
    </Toolbar>
  </MuiAppBar>
);