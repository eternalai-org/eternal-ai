from x_content.constants.utils import get_and_warn

CUSTOM_ANALYTIC_URL = get_and_warn("CUSTOM_ANALYTIC_URL").rstrip("/")
CUSTOM_ANALYTIC_API_KEY = get_and_warn("CUSTOM_ANALYTIC_API_KEY")
