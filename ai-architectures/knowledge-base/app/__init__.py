__version__ = 'v3.2.9'

import logging
import os

logger = logging.getLogger(__name__)

from dotenv import load_dotenv
if not load_dotenv(os.path.join(os.getcwd(), '.env')):
    logger.warning("No .env file found")

from . import (
    wrappers, 
)
