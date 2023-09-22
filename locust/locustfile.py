import api
import common

from locust import HttpUser, between, task


class ApiNestData(HttpUser):
    wait_time = between(2, 3)

    def on_start(self):
        print("on_start")
        self.api = api.Api(self.client)

    # @task
    # def get_root_path(self):
    #     self.api.get_text_root()

    # @task
    # def get_api_path(self):
    #     self.api.get_text_api()

    # @task
    # def call_plus(self):
    #     self.api.post_plus(common.get_random_int(), common.get_random_int())

    # @task
    # def call_minus(self):
    #     self.api.post_minus(common.get_random_int(), common.get_random_int())

    # @task
    # def call_multiple(self):
    #     self.api.post_multiple(common.get_random_int(),
    #                            common.get_random_int())

    # @task
    # def call_divide(self):
    #     self.api.post_divide(common.get_random_int(),
    #                          common.get_random_int())

    @task
    def call_product(self):
        self.api.post_product(str(common.get_random_int_product()))
