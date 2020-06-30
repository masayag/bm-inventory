# coding: utf-8

"""
    AssistedInstall

    Assisted installation  # noqa: E501

    OpenAPI spec version: 1.0.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from bm_inventory_client.models.host import Host  # noqa: F401,E501
from bm_inventory_client.models.host_network import HostNetwork  # noqa: F401,E501
from bm_inventory_client.models.image_info import ImageInfo  # noqa: F401,E501


class Cluster(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'kind': 'str',
        'id': 'str',
        'href': 'str',
        'name': 'str',
        'openshift_version': 'str',
        'image_info': 'ImageInfo',
        'base_dns_domain': 'str',
        'cluster_network_cidr': 'str',
        'cluster_network_host_prefix': 'int',
        'service_network_cidr': 'str',
        'api_vip': 'str',
        'machine_network_cidr': 'str',
        'ingress_vip': 'str',
        'ssh_public_key': 'str',
        'status': 'str',
        'status_info': 'str',
        'status_updated_at': 'datetime',
        'hosts': 'list[Host]',
        'updated_at': 'datetime',
        'created_at': 'datetime',
        'install_started_at': 'datetime',
        'install_completed_at': 'datetime',
        'host_networks': 'list[HostNetwork]',
        'pull_secret_set': 'bool',
        'ignition_generator_version': 'str'
    }

    attribute_map = {
        'kind': 'kind',
        'id': 'id',
        'href': 'href',
        'name': 'name',
        'openshift_version': 'openshift_version',
        'image_info': 'image_info',
        'base_dns_domain': 'base_dns_domain',
        'cluster_network_cidr': 'cluster_network_cidr',
        'cluster_network_host_prefix': 'cluster_network_host_prefix',
        'service_network_cidr': 'service_network_cidr',
        'api_vip': 'api_vip',
        'machine_network_cidr': 'machine_network_cidr',
        'ingress_vip': 'ingress_vip',
        'ssh_public_key': 'ssh_public_key',
        'status': 'status',
        'status_info': 'status_info',
        'status_updated_at': 'status_updated_at',
        'hosts': 'hosts',
        'updated_at': 'updated_at',
        'created_at': 'created_at',
        'install_started_at': 'install_started_at',
        'install_completed_at': 'install_completed_at',
        'host_networks': 'host_networks',
        'pull_secret_set': 'pull_secret_set',
        'ignition_generator_version': 'ignition_generator_version'
    }

    def __init__(self, kind=None, id=None, href=None, name=None, openshift_version=None, image_info=None, base_dns_domain=None, cluster_network_cidr=None, cluster_network_host_prefix=None, service_network_cidr=None, api_vip=None, machine_network_cidr=None, ingress_vip=None, ssh_public_key=None, status=None, status_info=None, status_updated_at=None, hosts=None, updated_at=None, created_at=None, install_started_at=None, install_completed_at=None, host_networks=None, pull_secret_set=None, ignition_generator_version=None):  # noqa: E501
        """Cluster - a model defined in Swagger"""  # noqa: E501

        self._kind = None
        self._id = None
        self._href = None
        self._name = None
        self._openshift_version = None
        self._image_info = None
        self._base_dns_domain = None
        self._cluster_network_cidr = None
        self._cluster_network_host_prefix = None
        self._service_network_cidr = None
        self._api_vip = None
        self._machine_network_cidr = None
        self._ingress_vip = None
        self._ssh_public_key = None
        self._status = None
        self._status_info = None
        self._status_updated_at = None
        self._hosts = None
        self._updated_at = None
        self._created_at = None
        self._install_started_at = None
        self._install_completed_at = None
        self._host_networks = None
        self._pull_secret_set = None
        self._ignition_generator_version = None
        self.discriminator = None

        self.kind = kind
        self.id = id
        self.href = href
        if name is not None:
            self.name = name
        if openshift_version is not None:
            self.openshift_version = openshift_version
        self.image_info = image_info
        if base_dns_domain is not None:
            self.base_dns_domain = base_dns_domain
        if cluster_network_cidr is not None:
            self.cluster_network_cidr = cluster_network_cidr
        if cluster_network_host_prefix is not None:
            self.cluster_network_host_prefix = cluster_network_host_prefix
        if service_network_cidr is not None:
            self.service_network_cidr = service_network_cidr
        if api_vip is not None:
            self.api_vip = api_vip
        if machine_network_cidr is not None:
            self.machine_network_cidr = machine_network_cidr
        if ingress_vip is not None:
            self.ingress_vip = ingress_vip
        if ssh_public_key is not None:
            self.ssh_public_key = ssh_public_key
        self.status = status
        self.status_info = status_info
        if status_updated_at is not None:
            self.status_updated_at = status_updated_at
        if hosts is not None:
            self.hosts = hosts
        if updated_at is not None:
            self.updated_at = updated_at
        if created_at is not None:
            self.created_at = created_at
        if install_started_at is not None:
            self.install_started_at = install_started_at
        if install_completed_at is not None:
            self.install_completed_at = install_completed_at
        if host_networks is not None:
            self.host_networks = host_networks
        if pull_secret_set is not None:
            self.pull_secret_set = pull_secret_set
        if ignition_generator_version is not None:
            self.ignition_generator_version = ignition_generator_version

    @property
    def kind(self):
        """Gets the kind of this Cluster.  # noqa: E501

        Indicates the type of this object. Will be 'Cluster' if this is a complete object or 'ClusterLink' if it is just a link.  # noqa: E501

        :return: The kind of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._kind

    @kind.setter
    def kind(self, kind):
        """Sets the kind of this Cluster.

        Indicates the type of this object. Will be 'Cluster' if this is a complete object or 'ClusterLink' if it is just a link.  # noqa: E501

        :param kind: The kind of this Cluster.  # noqa: E501
        :type: str
        """
        if kind is None:
            raise ValueError("Invalid value for `kind`, must not be `None`")  # noqa: E501
        allowed_values = ["Cluster"]  # noqa: E501
        if kind not in allowed_values:
            raise ValueError(
                "Invalid value for `kind` ({0}), must be one of {1}"  # noqa: E501
                .format(kind, allowed_values)
            )

        self._kind = kind

    @property
    def id(self):
        """Gets the id of this Cluster.  # noqa: E501

        Unique identifier of the object.  # noqa: E501

        :return: The id of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this Cluster.

        Unique identifier of the object.  # noqa: E501

        :param id: The id of this Cluster.  # noqa: E501
        :type: str
        """
        if id is None:
            raise ValueError("Invalid value for `id`, must not be `None`")  # noqa: E501

        self._id = id

    @property
    def href(self):
        """Gets the href of this Cluster.  # noqa: E501

        Self link.  # noqa: E501

        :return: The href of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._href

    @href.setter
    def href(self, href):
        """Sets the href of this Cluster.

        Self link.  # noqa: E501

        :param href: The href of this Cluster.  # noqa: E501
        :type: str
        """
        if href is None:
            raise ValueError("Invalid value for `href`, must not be `None`")  # noqa: E501

        self._href = href

    @property
    def name(self):
        """Gets the name of this Cluster.  # noqa: E501

        Name of the OpenShift cluster.  # noqa: E501

        :return: The name of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this Cluster.

        Name of the OpenShift cluster.  # noqa: E501

        :param name: The name of this Cluster.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def openshift_version(self):
        """Gets the openshift_version of this Cluster.  # noqa: E501

        Version of the OpenShift cluster.  # noqa: E501

        :return: The openshift_version of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._openshift_version

    @openshift_version.setter
    def openshift_version(self, openshift_version):
        """Sets the openshift_version of this Cluster.

        Version of the OpenShift cluster.  # noqa: E501

        :param openshift_version: The openshift_version of this Cluster.  # noqa: E501
        :type: str
        """
        allowed_values = ["4.5"]  # noqa: E501
        if openshift_version not in allowed_values:
            raise ValueError(
                "Invalid value for `openshift_version` ({0}), must be one of {1}"  # noqa: E501
                .format(openshift_version, allowed_values)
            )

        self._openshift_version = openshift_version

    @property
    def image_info(self):
        """Gets the image_info of this Cluster.  # noqa: E501


        :return: The image_info of this Cluster.  # noqa: E501
        :rtype: ImageInfo
        """
        return self._image_info

    @image_info.setter
    def image_info(self, image_info):
        """Sets the image_info of this Cluster.


        :param image_info: The image_info of this Cluster.  # noqa: E501
        :type: ImageInfo
        """
        if image_info is None:
            raise ValueError("Invalid value for `image_info`, must not be `None`")  # noqa: E501

        self._image_info = image_info

    @property
    def base_dns_domain(self):
        """Gets the base_dns_domain of this Cluster.  # noqa: E501

        Base domain of the cluster. All DNS records must be sub-domains of this base and include the cluster name.  # noqa: E501

        :return: The base_dns_domain of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._base_dns_domain

    @base_dns_domain.setter
    def base_dns_domain(self, base_dns_domain):
        """Sets the base_dns_domain of this Cluster.

        Base domain of the cluster. All DNS records must be sub-domains of this base and include the cluster name.  # noqa: E501

        :param base_dns_domain: The base_dns_domain of this Cluster.  # noqa: E501
        :type: str
        """

        self._base_dns_domain = base_dns_domain

    @property
    def cluster_network_cidr(self):
        """Gets the cluster_network_cidr of this Cluster.  # noqa: E501

        IP address block from which Pod IPs are allocated This block must not overlap with existing physical networks. These IP addresses are used for the Pod network, and if you need to access the Pods from an external network, configure load balancers and routers to manage the traffic.  # noqa: E501

        :return: The cluster_network_cidr of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._cluster_network_cidr

    @cluster_network_cidr.setter
    def cluster_network_cidr(self, cluster_network_cidr):
        """Sets the cluster_network_cidr of this Cluster.

        IP address block from which Pod IPs are allocated This block must not overlap with existing physical networks. These IP addresses are used for the Pod network, and if you need to access the Pods from an external network, configure load balancers and routers to manage the traffic.  # noqa: E501

        :param cluster_network_cidr: The cluster_network_cidr of this Cluster.  # noqa: E501
        :type: str
        """

        self._cluster_network_cidr = cluster_network_cidr

    @property
    def cluster_network_host_prefix(self):
        """Gets the cluster_network_host_prefix of this Cluster.  # noqa: E501

        The subnet prefix length to assign to each individual node. For example, if clusterNetworkHostPrefix is set to 23, then each node is assigned a /23 subnet out of the given cidr (clusterNetworkCIDR), which allows for 510 (2^(32 - 23) - 2) pod IPs addresses. If you are required to provide access to nodes from an external network, configure load balancers and routers to manage the traffic.  # noqa: E501

        :return: The cluster_network_host_prefix of this Cluster.  # noqa: E501
        :rtype: int
        """
        return self._cluster_network_host_prefix

    @cluster_network_host_prefix.setter
    def cluster_network_host_prefix(self, cluster_network_host_prefix):
        """Sets the cluster_network_host_prefix of this Cluster.

        The subnet prefix length to assign to each individual node. For example, if clusterNetworkHostPrefix is set to 23, then each node is assigned a /23 subnet out of the given cidr (clusterNetworkCIDR), which allows for 510 (2^(32 - 23) - 2) pod IPs addresses. If you are required to provide access to nodes from an external network, configure load balancers and routers to manage the traffic.  # noqa: E501

        :param cluster_network_host_prefix: The cluster_network_host_prefix of this Cluster.  # noqa: E501
        :type: int
        """
        if cluster_network_host_prefix is not None and cluster_network_host_prefix > 32:  # noqa: E501
            raise ValueError("Invalid value for `cluster_network_host_prefix`, must be a value less than or equal to `32`")  # noqa: E501
        if cluster_network_host_prefix is not None and cluster_network_host_prefix < 1:  # noqa: E501
            raise ValueError("Invalid value for `cluster_network_host_prefix`, must be a value greater than or equal to `1`")  # noqa: E501

        self._cluster_network_host_prefix = cluster_network_host_prefix

    @property
    def service_network_cidr(self):
        """Gets the service_network_cidr of this Cluster.  # noqa: E501

        The IP address pool to use for service IP addresses. You can enter only one IP address pool. If you need to access the services from an external network, configure load balancers and routers to manage the traffic.  # noqa: E501

        :return: The service_network_cidr of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._service_network_cidr

    @service_network_cidr.setter
    def service_network_cidr(self, service_network_cidr):
        """Sets the service_network_cidr of this Cluster.

        The IP address pool to use for service IP addresses. You can enter only one IP address pool. If you need to access the services from an external network, configure load balancers and routers to manage the traffic.  # noqa: E501

        :param service_network_cidr: The service_network_cidr of this Cluster.  # noqa: E501
        :type: str
        """

        self._service_network_cidr = service_network_cidr

    @property
    def api_vip(self):
        """Gets the api_vip of this Cluster.  # noqa: E501

        Virtual IP used to reach the OpenShift cluster API.  # noqa: E501

        :return: The api_vip of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._api_vip

    @api_vip.setter
    def api_vip(self, api_vip):
        """Sets the api_vip of this Cluster.

        Virtual IP used to reach the OpenShift cluster API.  # noqa: E501

        :param api_vip: The api_vip of this Cluster.  # noqa: E501
        :type: str
        """

        self._api_vip = api_vip

    @property
    def machine_network_cidr(self):
        """Gets the machine_network_cidr of this Cluster.  # noqa: E501

        A CIDR that all hosts belonging to the cluster should have an interfaces with IP address that belongs to this CIDR. The api_vip belongs to this CIDR.  # noqa: E501

        :return: The machine_network_cidr of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._machine_network_cidr

    @machine_network_cidr.setter
    def machine_network_cidr(self, machine_network_cidr):
        """Sets the machine_network_cidr of this Cluster.

        A CIDR that all hosts belonging to the cluster should have an interfaces with IP address that belongs to this CIDR. The api_vip belongs to this CIDR.  # noqa: E501

        :param machine_network_cidr: The machine_network_cidr of this Cluster.  # noqa: E501
        :type: str
        """

        self._machine_network_cidr = machine_network_cidr

    @property
    def ingress_vip(self):
        """Gets the ingress_vip of this Cluster.  # noqa: E501

        Virtual IP used for cluster ingress traffic.  # noqa: E501

        :return: The ingress_vip of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._ingress_vip

    @ingress_vip.setter
    def ingress_vip(self, ingress_vip):
        """Sets the ingress_vip of this Cluster.

        Virtual IP used for cluster ingress traffic.  # noqa: E501

        :param ingress_vip: The ingress_vip of this Cluster.  # noqa: E501
        :type: str
        """

        self._ingress_vip = ingress_vip

    @property
    def ssh_public_key(self):
        """Gets the ssh_public_key of this Cluster.  # noqa: E501

        SSH public key for debugging OpenShift nodes.  # noqa: E501

        :return: The ssh_public_key of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._ssh_public_key

    @ssh_public_key.setter
    def ssh_public_key(self, ssh_public_key):
        """Sets the ssh_public_key of this Cluster.

        SSH public key for debugging OpenShift nodes.  # noqa: E501

        :param ssh_public_key: The ssh_public_key of this Cluster.  # noqa: E501
        :type: str
        """

        self._ssh_public_key = ssh_public_key

    @property
    def status(self):
        """Gets the status of this Cluster.  # noqa: E501

        Status of the OpenShift cluster.  # noqa: E501

        :return: The status of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._status

    @status.setter
    def status(self, status):
        """Sets the status of this Cluster.

        Status of the OpenShift cluster.  # noqa: E501

        :param status: The status of this Cluster.  # noqa: E501
        :type: str
        """
        if status is None:
            raise ValueError("Invalid value for `status`, must not be `None`")  # noqa: E501
        allowed_values = ["insufficient", "ready", "error", "installing", "installed"]  # noqa: E501
        if status not in allowed_values:
            raise ValueError(
                "Invalid value for `status` ({0}), must be one of {1}"  # noqa: E501
                .format(status, allowed_values)
            )

        self._status = status

    @property
    def status_info(self):
        """Gets the status_info of this Cluster.  # noqa: E501

        Additional information pertaining to the status of the OpenShift cluster.  # noqa: E501

        :return: The status_info of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._status_info

    @status_info.setter
    def status_info(self, status_info):
        """Sets the status_info of this Cluster.

        Additional information pertaining to the status of the OpenShift cluster.  # noqa: E501

        :param status_info: The status_info of this Cluster.  # noqa: E501
        :type: str
        """
        if status_info is None:
            raise ValueError("Invalid value for `status_info`, must not be `None`")  # noqa: E501

        self._status_info = status_info

    @property
    def status_updated_at(self):
        """Gets the status_updated_at of this Cluster.  # noqa: E501

        The last time that the cluster status has been updated  # noqa: E501

        :return: The status_updated_at of this Cluster.  # noqa: E501
        :rtype: datetime
        """
        return self._status_updated_at

    @status_updated_at.setter
    def status_updated_at(self, status_updated_at):
        """Sets the status_updated_at of this Cluster.

        The last time that the cluster status has been updated  # noqa: E501

        :param status_updated_at: The status_updated_at of this Cluster.  # noqa: E501
        :type: datetime
        """

        self._status_updated_at = status_updated_at

    @property
    def hosts(self):
        """Gets the hosts of this Cluster.  # noqa: E501

        Hosts that are associated with this cluster.  # noqa: E501

        :return: The hosts of this Cluster.  # noqa: E501
        :rtype: list[Host]
        """
        return self._hosts

    @hosts.setter
    def hosts(self, hosts):
        """Sets the hosts of this Cluster.

        Hosts that are associated with this cluster.  # noqa: E501

        :param hosts: The hosts of this Cluster.  # noqa: E501
        :type: list[Host]
        """

        self._hosts = hosts

    @property
    def updated_at(self):
        """Gets the updated_at of this Cluster.  # noqa: E501

        The last time that this cluster was updated.  # noqa: E501

        :return: The updated_at of this Cluster.  # noqa: E501
        :rtype: datetime
        """
        return self._updated_at

    @updated_at.setter
    def updated_at(self, updated_at):
        """Sets the updated_at of this Cluster.

        The last time that this cluster was updated.  # noqa: E501

        :param updated_at: The updated_at of this Cluster.  # noqa: E501
        :type: datetime
        """

        self._updated_at = updated_at

    @property
    def created_at(self):
        """Gets the created_at of this Cluster.  # noqa: E501

        The time that this cluster was created.  # noqa: E501

        :return: The created_at of this Cluster.  # noqa: E501
        :rtype: datetime
        """
        return self._created_at

    @created_at.setter
    def created_at(self, created_at):
        """Sets the created_at of this Cluster.

        The time that this cluster was created.  # noqa: E501

        :param created_at: The created_at of this Cluster.  # noqa: E501
        :type: datetime
        """

        self._created_at = created_at

    @property
    def install_started_at(self):
        """Gets the install_started_at of this Cluster.  # noqa: E501

        The time that this cluster began installation.  # noqa: E501

        :return: The install_started_at of this Cluster.  # noqa: E501
        :rtype: datetime
        """
        return self._install_started_at

    @install_started_at.setter
    def install_started_at(self, install_started_at):
        """Sets the install_started_at of this Cluster.

        The time that this cluster began installation.  # noqa: E501

        :param install_started_at: The install_started_at of this Cluster.  # noqa: E501
        :type: datetime
        """

        self._install_started_at = install_started_at

    @property
    def install_completed_at(self):
        """Gets the install_completed_at of this Cluster.  # noqa: E501

        The time that this cluster completed installation.  # noqa: E501

        :return: The install_completed_at of this Cluster.  # noqa: E501
        :rtype: datetime
        """
        return self._install_completed_at

    @install_completed_at.setter
    def install_completed_at(self, install_completed_at):
        """Sets the install_completed_at of this Cluster.

        The time that this cluster completed installation.  # noqa: E501

        :param install_completed_at: The install_completed_at of this Cluster.  # noqa: E501
        :type: datetime
        """

        self._install_completed_at = install_completed_at

    @property
    def host_networks(self):
        """Gets the host_networks of this Cluster.  # noqa: E501

        List of host networks to be filled during query.  # noqa: E501

        :return: The host_networks of this Cluster.  # noqa: E501
        :rtype: list[HostNetwork]
        """
        return self._host_networks

    @host_networks.setter
    def host_networks(self, host_networks):
        """Sets the host_networks of this Cluster.

        List of host networks to be filled during query.  # noqa: E501

        :param host_networks: The host_networks of this Cluster.  # noqa: E501
        :type: list[HostNetwork]
        """

        self._host_networks = host_networks

    @property
    def pull_secret_set(self):
        """Gets the pull_secret_set of this Cluster.  # noqa: E501

        True if the pull-secret has been added to the cluster  # noqa: E501

        :return: The pull_secret_set of this Cluster.  # noqa: E501
        :rtype: bool
        """
        return self._pull_secret_set

    @pull_secret_set.setter
    def pull_secret_set(self, pull_secret_set):
        """Sets the pull_secret_set of this Cluster.

        True if the pull-secret has been added to the cluster  # noqa: E501

        :param pull_secret_set: The pull_secret_set of this Cluster.  # noqa: E501
        :type: bool
        """

        self._pull_secret_set = pull_secret_set

    @property
    def ignition_generator_version(self):
        """Gets the ignition_generator_version of this Cluster.  # noqa: E501


        :return: The ignition_generator_version of this Cluster.  # noqa: E501
        :rtype: str
        """
        return self._ignition_generator_version

    @ignition_generator_version.setter
    def ignition_generator_version(self, ignition_generator_version):
        """Sets the ignition_generator_version of this Cluster.


        :param ignition_generator_version: The ignition_generator_version of this Cluster.  # noqa: E501
        :type: str
        """

        self._ignition_generator_version = ignition_generator_version

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, Cluster):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
